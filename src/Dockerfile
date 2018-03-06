FROM golang:1.9.2-alpine3.7

RUN apk add --update tzdata bash wget curl git;

RUN mkdir -p $GOPATH/bin && \
    go get github.com/pilu/fresh && \
    go get github.com/gorilla/mux && \
    go get github.com/xeipuuv/gojsonschema
