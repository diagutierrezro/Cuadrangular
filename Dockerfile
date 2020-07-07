FROM golang:1.14.1-alpine

RUN apk add --no-cache git
RUN apk add --no-cache build-base

RUN set -x \
    && export GO111MODULE=on \
    && go get -u github.com/astaxie/beego \
    && go get -u github.com/beego/bee

WORKDIR /go/src/github.com/diagutierrezro/Cuadrangular

COPY . .
