#!/bin/bash

cd ui && yarn install --ignore-optional && yarn build && cd ..
go mod download

if [ "$1" = "--ci" ]; then
    # cross-compilation for CI
    GOOS=linux GOARCH=amd64 go build -o okr2go-linux-amd64 cmd/okr2go/okr2go.go
    GOOS=darwin GOARCH=arm64 go build -o okr2go-darwin-arm64 cmd/okr2go/okr2go.go
    GOOS=darwin GOARCH=amd64 go build -o okr2go-darwin-amd64 cmd/okr2go/okr2go.go
    GOOS=windows GOARCH=amd64 go build -o okr2go-windows-amd64.exe cmd/okr2go/okr2go.go
    go build cmd/okr2go/okr2go.go
fi

# install for local
go install cmd/okr2go/okr2go.go
