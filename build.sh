#!/bin/bash

set -e  # Exit if any command fails

APP_NAME="leetcode-helper"

echo "Tidying up..."
go mod tidy
go mod download

echo "Building $APP_NAME..."
go build -a -gcflags=all="-l -B" -ldflags="-w -s" -o "$APP_NAME"

echo "Compressing with UPX..."
upx --best --lzma "$APP_NAME"

echo "Build completed successfully!"
