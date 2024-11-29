module github.com/otel-warez/stdinotel/cmd/stdinotel

go 1.23.0

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/fileexporter v0.114.0
	github.com/open-telemetry/opentelemetry-collector-contrib/exporter/splunkhecexporter v0.114.0
	github.com/otel-warez/stdinreceiver v0.0.0-20241129045800-7ce5988f7d54
	go.opentelemetry.io/collector/component v0.114.0
	go.opentelemetry.io/collector/component/componentstatus v0.114.0
	go.opentelemetry.io/collector/config/configopaque v1.20.0
	go.opentelemetry.io/collector/config/configretry v1.20.0
	go.opentelemetry.io/collector/config/configtelemetry v0.114.0
	go.opentelemetry.io/collector/exporter v0.114.0
	go.opentelemetry.io/collector/exporter/otlpexporter v0.114.0
	go.opentelemetry.io/collector/exporter/otlphttpexporter v0.114.0
	go.opentelemetry.io/collector/pdata v1.20.0
	go.opentelemetry.io/collector/receiver v0.114.0
	go.opentelemetry.io/otel/metric v1.32.0
	go.opentelemetry.io/otel/sdk v1.32.0
	go.uber.org/zap v1.27.0
)

require (
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/goccy/go-json v0.10.3 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.2 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mostynb/go-grpc-compression v1.2.3 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.114.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/sharedcomponent v0.114.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk v0.114.0 // indirect
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/batchperresourceattr v0.114.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rogpeppe/go-internal v1.13.1 // indirect
	github.com/rs/cors v1.11.1 // indirect
	github.com/stretchr/testify v1.9.0 // indirect
	go.opentelemetry.io/collector v0.114.0 // indirect
	go.opentelemetry.io/collector/client v1.20.0 // indirect
	go.opentelemetry.io/collector/config/configauth v0.114.0 // indirect
	go.opentelemetry.io/collector/config/configcompression v1.20.0 // indirect
	go.opentelemetry.io/collector/config/configgrpc v0.114.0 // indirect
	go.opentelemetry.io/collector/config/confighttp v0.114.0 // indirect
	go.opentelemetry.io/collector/config/confignet v1.20.0 // indirect
	go.opentelemetry.io/collector/config/configtls v1.20.0 // indirect
	go.opentelemetry.io/collector/config/internal v0.114.0 // indirect
	go.opentelemetry.io/collector/confmap v1.20.0 // indirect
	go.opentelemetry.io/collector/consumer v0.114.0 // indirect
	go.opentelemetry.io/collector/consumer/consumererror v0.114.0 // indirect
	go.opentelemetry.io/collector/consumer/consumererror/consumererrorprofiles v0.114.0 // indirect
	go.opentelemetry.io/collector/consumer/consumerprofiles v0.114.0 // indirect
	go.opentelemetry.io/collector/exporter/exporterhelper/exporterhelperprofiles v0.114.0 // indirect
	go.opentelemetry.io/collector/exporter/exporterprofiles v0.114.0 // indirect
	go.opentelemetry.io/collector/extension v0.114.0 // indirect
	go.opentelemetry.io/collector/extension/auth v0.114.0 // indirect
	go.opentelemetry.io/collector/extension/experimental/storage v0.114.0 // indirect
	go.opentelemetry.io/collector/pdata/pprofile v0.114.0 // indirect
	go.opentelemetry.io/collector/pipeline v0.114.0 // indirect
	go.opentelemetry.io/collector/pipeline/pipelineprofiles v0.114.0 // indirect
	go.opentelemetry.io/collector/semconv v0.114.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.56.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.56.0 // indirect
	go.opentelemetry.io/otel v1.32.0 // indirect
	go.opentelemetry.io/otel/trace v1.32.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.31.0 // indirect
	golang.org/x/sys v0.27.0 // indirect
	golang.org/x/term v0.26.0 // indirect
	golang.org/x/text v0.20.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241104194629-dd2ea8efbc28 // indirect
	google.golang.org/grpc v1.67.1 // indirect
	google.golang.org/protobuf v1.35.1 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
