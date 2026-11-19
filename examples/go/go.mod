module github.com/traust-security/traust-sdk/examples/go

go 1.26

require (
	github.com/oapi-codegen/runtime v1.8.0
	github.com/traust-security/traust-sdk/go v0.0.0
)

require (
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/santhosh-tekuri/jsonschema/v6 v6.0.3 // indirect
	golang.org/x/text v0.39.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/traust-security/traust-sdk/go => ../../go
