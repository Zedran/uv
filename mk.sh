#!/usr/bin/env bash

go mod tidy
go build -trimpath -ldflags "-s -w" -o build/uv ./src
