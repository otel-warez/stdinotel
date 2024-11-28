// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stdinreceiver

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"go.opentelemetry.io/collector/component/componentstatus"
	"golang.org/x/term"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
)

var (
	stdin = os.Stdin
)

// stdinReceiver implements the component.MetricsReceiver for stdin metric protocol.
type stdinReceiver struct {
	config       *Config
	logsConsumer consumer.Logs
	obsrecv      *receiverhelper.ObsReport
	done         chan bool
}

// newLogsReceiver creates the stdin receiver with the given configuration.
func newLogsReceiver(
	_ context.Context,
	settings receiver.Settings,
	config component.Config,
	nextConsumer consumer.Logs,
) (receiver.Logs, error) {

	cfg := config.(*Config)

	obsrecv, err := receiverhelper.NewObsReport(receiverhelper.ObsReportSettings{
		ReceiverID:             settings.ID,
		Transport:              "",
		ReceiverCreateSettings: settings,
	})
	if err != nil {
		return nil, err
	}

	r := &stdinReceiver{
		config:       cfg,
		logsConsumer: nextConsumer,
		obsrecv:      obsrecv,
		done:         make(chan bool, 1),
	}

	return r, nil
}

func (r *stdinReceiver) runStdinInteractive(ctx context.Context, host component.Host) {
	reader := bufio.NewReader(stdin)
	scanner := bufio.NewScanner(reader)
	if term.IsTerminal(int(os.Stdout.Fd())) {
		interruptChan := make(chan os.Signal, 1)
		go func() {
			select {
			case <-interruptChan:
				componentstatus.ReportStatus(host, componentstatus.NewRecoverableErrorEvent(errors.New("stdin interrupt")))
			case <-r.done:
				return
			}
		}()
		signal.Notify(interruptChan, syscall.SIGINT, syscall.SIGTERM)
	}

	scanner.Split(bufio.ScanLines) // Set up the split function.

	for {
		switch {
		case scanner.Scan():
			if scanner.Err() != nil {
				componentstatus.ReportStatus(host, componentstatus.NewRecoverableErrorEvent(errors.New("stdin closed")))
				continue
			}
			line := scanner.Text()
			if line == "" {
				componentstatus.ReportStatus(host, componentstatus.NewRecoverableErrorEvent(errors.New("user end of input")))
				continue
			}
			r.obsrecv.StartLogsOp(ctx)
			err := r.consumeLine(ctx, line)
			r.obsrecv.EndLogsOp(ctx, "stdin", 1, err)
			if err != nil {
				componentstatus.ReportStatus(host, componentstatus.NewPermanentErrorEvent(err))
			}
		case <-r.done:
			return
		}

	}
}

func (r *stdinReceiver) runStdinPiped(ctx context.Context, host component.Host) {
	var scanner *bufio.Scanner
	data, err := io.ReadAll(stdin)
	if err != nil {
		componentstatus.ReportStatus(host, componentstatus.NewFatalErrorEvent(errors.New("cannot read stdin")))
	}
	scanner = bufio.NewScanner(bytes.NewReader(data))

	scanner.Split(bufio.ScanLines) // Set up the split function.

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			componentstatus.ReportStatus(host, componentstatus.NewPermanentErrorEvent(err))
			continue
		}
		line := scanner.Text()
		r.obsrecv.StartLogsOp(ctx)
		err := r.consumeLine(ctx, line)
		r.obsrecv.EndLogsOp(ctx, "stdin", 1, err)
		if err != nil {
			componentstatus.ReportStatus(host, componentstatus.NewPermanentErrorEvent(err))
		}
	}
	componentstatus.ReportStatus(host, componentstatus.NewEvent(componentstatus.StatusStopping))
}

// Start starts the stdin receiver.
func (r *stdinReceiver) Start(ctx context.Context, host component.Host) error {
	if isInputPiped() {
		go r.runStdinPiped(ctx, host)
	} else {
		go r.runStdinInteractive(ctx, host)
	}
	return nil
}

func (r *stdinReceiver) consumeLine(ctx context.Context, line string) error {
	ld := plog.NewLogs()
	rl := ld.ResourceLogs().AppendEmpty()
	sl := rl.ScopeLogs().AppendEmpty()
	lr := sl.LogRecords().AppendEmpty()
	lr.Body().SetStr(line)
	lr.SetTimestamp(pcommon.NewTimestampFromTime(time.Now()))
	err := r.logsConsumer.ConsumeLogs(ctx, ld)
	return err
}

// Shutdown shuts down the stdin receiver.
func (r *stdinReceiver) Shutdown(context.Context) error {
	if r.done != nil {
		close(r.done)
	}
	return nil
}

func isInputPiped() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}
