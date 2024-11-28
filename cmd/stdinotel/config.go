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

package main

import (
	"context"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/fileexporter"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/exporter"
	"os"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/splunkhecexporter"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/exporter/otlpexporter"
	"go.opentelemetry.io/collector/exporter/otlphttpexporter"
)

func createExporter(settings component.TelemetrySettings) (exporter.Logs, error) {
	protocol := getProtocol()
	exporters, _ := exporter.MakeFactoryMap(
		otlpexporter.NewFactory(),
		otlphttpexporter.NewFactory(),
		fileexporter.NewFactory(),
		splunkhecexporter.NewFactory(),
	)

	cfg := exporters[protocol].CreateDefaultConfig()

	switch cfg.(type) {
	case *splunkhecexporter.Config:
		splunkCfg := cfg.(*splunkhecexporter.Config)
		if endpoint := os.Getenv("STDINOTEL_ENDPOINT"); endpoint != "" {
			splunkCfg.Endpoint = endpoint
		}
		splunkCfg.Token = configopaque.String(os.Getenv("STDINOTEL_TOKEN"))
		splunkCfg.Index = os.Getenv("STDINOTEL_SPLUNK_INDEX")
		splunkCfg.TLSSetting.InsecureSkipVerify = os.Getenv("STDINOTEL_TLS_INSECURE_SKIP_VERIFY") == "true"
		splunkCfg.BackOffConfig = configretry.BackOffConfig{
			Enabled: false,
		}
	case *otlpexporter.Config:
		otlpCfg := cfg.(*otlpexporter.Config)
		if endpoint := os.Getenv("STDINOTEL_ENDPOINT"); endpoint != "" {
			otlpCfg.Endpoint = endpoint
		} else {
			otlpCfg.Endpoint = "dns:///localhost:4317"
		}
		if os.Getenv("STDINOTEL_TLS_INSECURE_SKIP_VERIFY") == "true" {
			otlpCfg.TLSSetting.InsecureSkipVerify = true
			otlpCfg.TLSSetting.Insecure = true
		}
		otlpCfg.RetryConfig = configretry.BackOffConfig{
			Enabled: false,
		}
	case *otlphttpexporter.Config:
		otlpCfg := cfg.(*otlphttpexporter.Config)
		if endpoint := os.Getenv("STDINOTEL_ENDPOINT"); endpoint != "" {
			otlpCfg.Endpoint = endpoint
		} else {
			otlpCfg.Endpoint = "dns:///localhost:4318"
		}
		if os.Getenv("STDINOTEL_TLS_INSECURE_SKIP_VERIFY") == "true" {
			otlpCfg.TLSSetting.InsecureSkipVerify = true
			otlpCfg.TLSSetting.Insecure = true
		}
		otlpCfg.RetryConfig = configretry.BackOffConfig{
			Enabled: false,
		}
	}

	return exporters[protocol].CreateLogs(context.Background(), exporter.Settings{
		TelemetrySettings: settings,
	}, cfg)
}

func getProtocol() component.Type {
	protocol := os.Getenv("STDINOTEL_PROTOCOL")
	switch protocol {
	case "splunk_hec":
		return component.MustNewType("splunk_hec")
	case "otlp":
		return component.MustNewType("otlp")
	case "otlphttp":
		return component.MustNewType("otlphttp")
	default:
		return component.MustNewType("otlp")
	}
}
