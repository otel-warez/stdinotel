// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Program stdinotel is an OpenTelemetry Collector binary.
package main

import (
	"context"
	"log"

	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/otel/sdk/trace"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componentstatus"
	"go.opentelemetry.io/collector/config/configtelemetry"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/otel/metric"
	noopmetric "go.opentelemetry.io/otel/metric/noop"
	"go.uber.org/zap"

	"github.com/otel-warez/stdinreceiver"
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
