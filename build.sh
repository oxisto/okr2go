#!/bin/bash

cd ui && yarn install --ignore-optional && yarn build && cd ..
go mod download

if [ "$1" = "--ci" ]; then
    # cross-compilation for CI
    GOOS=linux GOARCH=arm64 go build cmd/okr2go/okr2go.go && zip okr2go-linux-arm64.zip okr2go
    GOOS=linux GOARCH=amd64 go build cmd/okr2go/okr2go.go && zip okr2go-linux-amd64.zip okr2go
    GOOS=darwin GOARCH=arm64 go build cmd/okr2go/okr2go.go && zip okr2go-darwin-arm64.zip okr2go
    GOOS=darwin GOARCH=amd64 go build cmd/okr2go/okr2go.go && zip okr2go-darwin-amd64.zip okr2go
    GOOS=windows GOARCH=amd64 go build cmd/okr2go/okr2go.go && zip okr2go-windows-amd64.zip okr2go.exe
fi

go build cmd/okr2go/okr2go.go

# install for local
go install cmd/okr2go/okr2go.go
