#!/usr/bin/env bash

set -Eeuo pipefail

go build -o "$GOPATH/bin/go-generate-struct-interface" ./cmd/go-generate-struct-interface
