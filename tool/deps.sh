#!/bin/bash

source tool/env.sh

# Download dependencies
echo ""
echo ">>> Downloading dependencies"
echo ""

export GOPATH=$PWD/.vendor

rm -rf $PWD/.vendor
mkdir $PWD/.vendor

go get -u -v \
	github.com/go-martini/martini \
	github.com/codegangsta/gin