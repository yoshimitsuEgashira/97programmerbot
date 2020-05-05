FROM golang:1.12.14-alpine3.10

ENV GOPATH $GOPATH:/go/src
WORKDIR /go/src/github.com/yoshimitsuEgashira/97programmerbot

COPY . /go/src/github.com/yoshimitsuEgashira/97programmerbot

ENV GO111MODULE=on

RUN set -eux && \
    apk update && \
    apk add --no-cache git && \
    go fmt ./... && \
    go get -u

CMD ["go", "run", "./main.go"]
