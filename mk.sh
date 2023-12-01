#!/usr/bin/env bash

go test -v ./...
go build -trimpath -ldflags "-s -w" -o build/uv ./src
