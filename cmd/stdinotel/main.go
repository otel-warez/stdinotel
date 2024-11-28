// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Program stdinotel is an OpenTelemetry Collector binary.
package main

import (
	"context"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/otel/sdk/trace"
	"log"

	"github.com/otel-warez/stdinotel/receiver/stdinreceiver"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componentstatus"
	"go.opentelemetry.io/collector/config/configtelemetry"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/otel/metric"
	noopmetric "go.opentelemetry.io/otel/metric/noop"
	"go.uber.org/zap"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	h := ttyHost{
		errStatus: make(chan error, 1),
	}
	var logger *zap.Logger
	logger, err := zap.NewDevelopment()
	settings := component.TelemetrySettings{
		Logger: logger,
		LeveledMeterProvider: func(_ configtelemetry.Level) metric.MeterProvider {
			return noopmetric.NewMeterProvider()
		},
		TracerProvider: trace.NewTracerProvider(),
		MeterProvider:  noopmetric.NewMeterProvider(),
		MetricsLevel:   configtelemetry.LevelNone,
		Resource:       pcommon.NewResource(),
	}

	e, err := createExporter(settings)
	if err != nil {
		return err
	}

	ctx := context.Background()

	r, err := stdinreceiver.NewFactory().CreateLogs(ctx, receiver.Settings{
		TelemetrySettings: settings,
	}, &stdinreceiver.Config{}, e)
	if err != nil {
		return err
	}

	err = e.Start(ctx, h)
	if err != nil {
		return err
	}
	err = r.Start(ctx, h)
	if err != nil {
		return err
	}

	err = <-h.errStatus

	_ = r.Shutdown(ctx)
	_ = e.Shutdown(ctx)

	return err
}

var _ component.Host = ttyHost{}
var _ componentstatus.Reporter = ttyHost{}

type ttyHost struct {
	errStatus chan error
}

func (t ttyHost) Report(event *componentstatus.Event) {
	if event.Status() == componentstatus.StatusStopping {
		close(t.errStatus)
	}
	if event.Err() != nil {
		t.errStatus <- event.Err()
	}
}

func (t ttyHost) GetExtensions() map[component.ID]component.Component {
	return nil
}
