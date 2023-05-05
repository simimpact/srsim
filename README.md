# srsim
A certain gacha game simulator


## Design notes

`pipeline`:
- data extraction

`protos`:
- model definition

`pkg/core`:
- core interfaces; glues everything together

`pkg/simulator`:
- 

`internal/character`:
- character implementation

`internal/cones`:
- cones/weapon implementation

`internal/relic`:
- relic implementation

`internal/enemy`:
- enemy implementation


## Dev environment requirements:

You will need to have the following installed:
- go
- node
- yarn
- protoc (Currently on: https://github.com/protocolbuffers/protobuf/releases/tag/v21.12)
- protoc-gen-go (Currently on v 1.30.0: `go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30.0`)
- protoc-gen-go-grpc (Currently on v1.3.0 `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3`)