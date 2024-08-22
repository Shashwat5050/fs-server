module iceline-hosting.com/fs-server

go 1.20

replace iceline-hosting.com/backend/proto => ../proto

replace iceline-hosting.com/core => ../core

require (
	github.com/pkg/sftp v1.13.6
	github.com/stretchr/testify v1.8.1
	github.com/uptrace/opentelemetry-go-extra/otelzap v0.1.17
	go.uber.org/zap v1.24.0
	golang.org/x/crypto v0.21.0
	golang.org/x/sync v0.6.0
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.34.2
	iceline-hosting.com/backend/proto v0.0.0-00010101000000-000000000000
	iceline-hosting.com/core v0.0.0-00010101000000-000000000000
)

require (
	github.com/aws/aws-sdk-go v1.34.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/kr/fs v0.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/uptrace/opentelemetry-go-extra/otelutil v0.1.17 // indirect
	go.opentelemetry.io/otel v1.11.1 // indirect
	go.opentelemetry.io/otel/trace v1.11.1 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto v0.0.0-20230306155012-7f2fa6fef1f4 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
