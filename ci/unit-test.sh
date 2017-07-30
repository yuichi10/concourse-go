#!/bin/bash

set -eu

dir="/go/src/github.com/yuichi10/pcf-app"
mkdir -p $dir
cp -r ./* $dir
pushd $dir
go test ./...
popd