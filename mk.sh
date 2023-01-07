#!/usr/bin/env bash

go test -v ./src
go build -trimpath -ldflags "-s -w" -o build/uv ./src
