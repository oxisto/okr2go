#!/bin/bash

cd okr2go-ui && yarn install --ignore-optional && yarn build --prod && cd ..
go mod download
packr build