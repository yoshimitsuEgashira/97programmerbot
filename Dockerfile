FROM golang:1.14-alpine AS build

LABEL maintainer="yossiee <ysmtegsr@gmail.com>" \
      image.registry="https://hub.docker.com/r/yossiee/97programmerbot" \
      image.source="https://github.com/yossiee/97programmerbot"

ENV GO111MODULE on

WORKDIR /go/src/97programmerbot
COPY . /go/src/97programmerbot

RUN set -eux && \
    apk update && \
    apk add --no-cache git && \
    go build -o app /bin/app

FROM scratch

COPY --from=build /bin/app /bin/app

ENTRYPOINT [ "/bin/app" ]
