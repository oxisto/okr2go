#!/bin/bash

cd okr2go-ui && yarn install --ignore-optional && yarn build --prod --no-progress && cd ..
go mod download
packr build cmd/okr2go/okr2go.go
