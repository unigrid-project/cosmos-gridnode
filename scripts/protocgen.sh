#!/usr/bin/env bash

set -e

echo "Generating proto code"
cd ../proto

buf generate --template buf.gen.yaml

cd ..

go mod tidy
