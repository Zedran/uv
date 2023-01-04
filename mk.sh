#!/usr/bin/env bash

go build -trimpath -ldflags "-s -w" -o build/uv ./src
