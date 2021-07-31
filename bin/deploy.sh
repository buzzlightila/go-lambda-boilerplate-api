#!/bin/sh

set -e

echo "Building..." >&2
export CGO_ENABLED=0
go mod download
go build \
    -trimpath \
    -o main \
    cmd/* && zip main.zip ./main