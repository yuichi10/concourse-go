#!/bin/bash

set -eu

dir="/go/src/github.com/yuichi10/concourse-go"
mkdir -p $dir
cp -r ./* $dir
pushd $dir
go test ./...
popd