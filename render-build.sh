#!/usr/bin/env bash
# exit on error
set -o errexit

go build -tags netgo -ldflags '-s -w' -o yoinker ./main.go