#!/bin/bash

set -eu
shopt -s dotglob

# glide install
cp -r ./* $PWD/$1