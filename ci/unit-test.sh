#!/bin/bash

set -eu

cp ./* /go/src/
pushd /go/src/
go test ./...
popd