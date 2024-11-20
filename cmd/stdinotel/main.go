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
	"github.com/otel-warez/stdinotel/receiver/stdinreceiver"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/exporter/exportertest"
	"go.opentelemetry.io/collector/otelcol"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"log"
)

func main() {
	factories, err := components()
	if err != nil {
		log.Fatalf("failed to build components: %v", err)
	}

	if err := run(context.Background(), factories); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, factories otelcol.Factories) error {
	cfg, id := createExporterConfig(factories)
	exporterFactory := factories.Exporters[id]
	e, err := exporterFactory.CreateLogs(ctx, exportertest.NewNopSettings(), cfg)
	if err != nil {
		return err
	}

	waitCh := make(chan struct{})
	r, err := factories.Receivers[component.MustNewType("stdin")].CreateLogs(ctx, receivertest.NewNopSettings(), &stdinreceiver.Config{
		StdinClosedHook: func() {
			<-waitCh
		},
	}, e)
	if err != nil {
		return err
	}
	err = e.Start(ctx, componenttest.NewNopHost())
	if err != nil {
		return err
	}
	err = r.Start(ctx, componenttest.NewNopHost())
	if err != nil {
		return err
	}
	<-waitCh

	_ = r.Shutdown(ctx)
	return e.Shutdown(ctx)
}
