module github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter

go 1.21

require (
	github.com/aws/aws-sdk-go v1.50.27
	github.com/google/uuid v1.6.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/awsutil v0.96.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/cwlogs v0.96.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/metrics v0.96.0
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal v0.96.0
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/resourcetotelemetry v0.96.0
	github.com/stretchr/testify v1.9.0
	go.opentelemetry.io/collector/component v0.96.1-0.20240315132530-eb5d2b9fbd12
	go.opentelemetry.io/collector/confmap v0.96.1-0.20240315132530-eb5d2b9fbd12
	go.opentelemetry.io/collector/consumer v0.96.1-0.20240315132530-eb5d2b9fbd12
	go.opentelemetry.io/collector/exporter v0.96.1-0.20240315132530-eb5d2b9fbd12
	go.opentelemetry.io/collector/featuregate v1.3.1-0.20240315165441-06f177a66b26
	go.opentelemetry.io/collector/pdata v1.3.1-0.20240315165441-06f177a66b26
	go.opentelemetry.io/collector/semconv v0.96.1-0.20240315132530-eb5d2b9fbd12
	go.opentelemetry.io/otel/metric v1.24.0
	go.opentelemetry.io/otel/trace v1.24.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.27.0
	golang.org/x/exp v0.0.0-20240103183307-be819d1f06fc
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.0.0-alpha.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.19.0 // indirect
	github.com/prometheus/client_model v0.6.0 // indirect
	github.com/prometheus/common v0.48.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	go.opentelemetry.io/collector v0.96.1-0.20240315132530-eb5d2b9fbd12 // indirect
	go.opentelemetry.io/collector/config/configretry v0.96.1-0.20240315132530-eb5d2b9fbd12 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.96.1-0.20240315132530-eb5d2b9fbd12 // indirect
	go.opentelemetry.io/collector/extension v0.96.1-0.20240315132530-eb5d2b9fbd12 // indirect
	go.opentelemetry.io/collector/receiver v0.96.1-0.20240315132530-eb5d2b9fbd12 // indirect
	go.opentelemetry.io/otel v1.24.0 // indirect
	go.opentelemetry.io/otel/exporters/prometheus v0.46.0 // indirect
	go.opentelemetry.io/otel/sdk v1.24.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.24.0 // indirect
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240123012728-ef4313101c80 // indirect
	google.golang.org/grpc v1.62.1 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/metrics => ../../internal/aws/metrics

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/awsutil => ../../internal/aws/awsutil

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/cwlogs => ../../internal/aws/cwlogs

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal => ../../internal/coreinternal

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/resourcetotelemetry => ../../pkg/resourcetotelemetry

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest => ../../pkg/pdatatest

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatautil => ../../pkg/pdatautil

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/golden => ../../pkg/golden
