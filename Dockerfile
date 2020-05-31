FROM golang:1.14-alpine

LABEL maintainer="yossiee <ysmtegsr@gmail.com>" \
    image.registry="https://hub.docker.com/r/yossiee/97programmer" \
    image.source="https://github.com/yossiee/97programmer"

ENV CGO_ENABLED=0

WORKDIR /go/src/app
COPY . /go/src/app

RUN set -eux && \
    apk update && \
    apk add --no-cache git

CMD ["go", "run", "main.go"]
