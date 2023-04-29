//go:generate sh -c "protoc --experimental_allow_proto3_optional --go_out=module=github.com/simimpact/srsim:. --go-grpc_out=module=github.com/simimpact/srsim:. pb/**/*.proto"
package srsim
