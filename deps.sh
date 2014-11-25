#!/bin/bash

# Download dependencies
echo ""
echo ">>> Downloading dependencies"
echo ""

export GOPATH=$(pwd)/.vendor

rm -rf .vendor
mkdir .vendor

go get -u -v \
	github.com/go-martini/martini \
	github.com/codegangsta/gin