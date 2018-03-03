#!/bin/bash
docker run -it --rm -v "$PWD":/go/src/tuliospuri/go-clean -p 80:8084 go-clean "$@"
