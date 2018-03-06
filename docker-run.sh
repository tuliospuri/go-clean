#!/bin/bash
docker run -it --rm -v "$PWD":/go/src/tuliospuri/go-clean -w /go/src/tuliospuri/go-clean/src  -p 80:8084 go-clean "$@"
