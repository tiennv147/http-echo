#!/usr/bin/env bash
set -euo pipefail

BACKEND_OUT="$GOPATH/src"
PROTO_FILE="./pb/config.proto"
PROTO_GEN_VALIDATE_LIB="$GOPATH/src/github.com/envoyproxy/protoc-gen-validate"

protoc -I. -I"${PROTO_GEN_VALIDATE_LIB}" \
--go_out="$BACKEND_OUT" \
--validate_out="lang=go:$BACKEND_OUT" \
"${PROTO_FILE}"