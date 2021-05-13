#!/bin/bash

cd ui && yarn install --ignore-optional && yarn build && cd ..
go mod download

# cross-compilation for CI
GOOS=linux GOARCH=amd64 go build -o bin/linux-amd64/okr2go cmd/okr2go/okr2go.go
GOOS=darwin GOARCH=arm64 go build -o bin/darwin-arm64/okr2go cmd/okr2go/okr2go.go
GOOS=darwin GOARCH=amd64 go build -o bin/darwin-amd64/okr2go cmd/okr2go/okr2go.go
GOOS=windows GOARCH=amd64 go build -o bin/windows-amd64/okr2go.exe cmd/okr2go/okr2go.go

# install for local
go install cmd/okr2go/okr2go.go
