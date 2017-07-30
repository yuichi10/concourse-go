#!/bin/bash

set -eu

mv ./* /go/src/
go test /go/src/...