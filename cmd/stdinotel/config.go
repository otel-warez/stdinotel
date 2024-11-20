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
	"os"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/splunkhecexporter"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/exporter/otlpexporter"
	"go.opentelemetry.io/collector/exporter/otlphttpexporter"
	"go.opentelemetry.io/collector/otelcol"
)

func createExporterConfig(factories otelcol.Factories) (component.Config, component.Type) {
	protocol := getProtocol()
	cfg := factories.Exporters[protocol].CreateDefaultConfig()
	if splunkCfg, ok := cfg.(*splunkhecexporter.Config); ok {
		splunkCfg.Endpoint = os.Getenv("STDINOTEL_ENDPOINT")
		splunkCfg.Token = configopaque.String(os.Getenv("STDINOTEL_TOKEN"))
		splunkCfg.Index = os.Getenv("STDINOTEL_SPLUNK_INDEX")
		splunkCfg.TLSSetting.InsecureSkipVerify = os.Getenv("STDINOTEL_TLS_INSECURE_SKIP_VERIFY") == "true"

	}
	if otlpCfg, ok := cfg.(*otlpexporter.Config); ok {
		otlpCfg.Endpoint = os.Getenv("STDINOTEL_ENDPOINT")
		otlpCfg.TLSSetting.InsecureSkipVerify = os.Getenv("STDINOTEL_TLS_INSECURE_SKIP_VERIFY") == "true"
	}
	if otlphttpCfg, ok := cfg.(*otlphttpexporter.Config); ok {
		otlphttpCfg.Endpoint = os.Getenv("STDINOTEL_ENDPOINT")
		otlphttpCfg.TLSSetting.InsecureSkipVerify = os.Getenv("STDINOTEL_TLS_INSECURE_SKIP_VERIFY") == "true"
	}

	return cfg, protocol
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
