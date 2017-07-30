#!/bin/bash

set -eu

cp -r ./* /go/src/
pushd /go/src/
go test ./...
popd