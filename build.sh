#!/bin/bash

cd ui && yarn install --ignore-optional && yarn build && cd ..
go mod download

if [ "$1" = "--ci" ]; then
    # cross-compilation for CI
    GOOS=linux GOARCH=amd64 go build cmd/okr2go/okr2go.go && zip okr2go-linux-x86_64.zip okr2go
    GOOS=darwin GOARCH=arm64 go build cmd/okr2go/okr2go.go && zip okr2go-macos-aarch_64.zip okr2go
    GOOS=darwin GOARCH=amd64 go build cmd/okr2go/okr2go.go && zip okr2go-macos-x86_64.zip okr2go
    GOOS=windows GOARCH=amd64 go build cmd/okr2go/okr2go.go && zip okr2go-win64.zip okr2go.exe
fi

go build cmd/okr2go/okr2go.go

# install for local
go install cmd/okr2go/okr2go.go
