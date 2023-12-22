#!/bin/bash

# sudo apt-get install gcc-multilib
# sudo apt-get install gcc-mingw-w64

# notice how we avoid spaces in $now to avoid quotation hell in go build command
now=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
GOOS=windows GOARCH=amd64 go build -ldflags "-X main.sha1ver=`git rev-parse HEAD` -X main.buildTime=$now" -o srsim_windows_amd64.exe ./cmd/srsim
GOOS=darwin GOARCH=arm64 go build -ldflags "-X main.sha1ver=`git rev-parse HEAD` -X main.buildTime=$now" -o srsim_darwin_arm64 ./cmd/srsim 
GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.sha1ver=`git rev-parse HEAD` -X main.buildTime=$now" -o srsim_darwin_amd64 ./cmd/srsim 
GOOS=linux GOARCH=amd64 go build -ldflags "-X main.sha1ver=`git rev-parse HEAD` -X main.buildTime=$now" -o srsim_linux_amd64 ./cmd/srsim 