#!/bin/bash

set -eu

dir="/go/src/github.com/yuichi10/concourse-go"
mkdir -p $dir
cp -r ./* $dir
pushd /go/src/
go test ./...
popd