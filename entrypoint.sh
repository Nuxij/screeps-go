#!/bin/bash
export PATH=/root/go/bin:$PATH
export GOPHERJS_GOROOT="$(go1.12.16 env GOROOT)"
export GOPATH=$(pwd)

exec "$@"
