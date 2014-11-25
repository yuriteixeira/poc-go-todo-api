#!/bin/bash

# Set needed env variables
export GOPATH=$(pwd)/.vendor:$(pwd)
export PATH=$(pwd)/.vendor/bin:$(pwd)/bin:$PATH
