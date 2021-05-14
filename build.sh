#!/bin/bash

cd ui && yarn install --ignore-optional && yarn build && cd ..
go mod download

if [ "$1" = "--ci" ]; then
    mkdir -p build

    # cross-compilation for CI
    GOOS=linux GOARCH=arm64 go build cmd/okr2go/okr2go.go && tar -czf build/okr2go-linux-arm64.tar.gz okr2go
    GOOS=linux GOARCH=amd64 go build cmd/okr2go/okr2go.go && tar -czf build/okr2go-linux-amd64.tar.gz okr2go
    GOOS=darwin GOARCH=arm64 go build cmd/okr2go/okr2go.go && tar -czf build/okr2go-darwin-arm64.tar.gz okr2go
    GOOS=darwin GOARCH=amd64 go build cmd/okr2go/okr2go.go && tar -czf build/okr2go-darwin-amd64.tar.gz okr2go
    GOOS=windows GOARCH=amd64 go build cmd/okr2go/okr2go.go && tar -czf build/okr2go-windows-amd64.tar.gz okr2go.exe
fi

go build cmd/okr2go/okr2go.go

# install for local
go install cmd/okr2go/okr2go.go
