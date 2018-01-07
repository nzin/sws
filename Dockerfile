FROM golang:1.9-alpine3.7

RUN apk add --no-cache \
    git \
    gcc \
    cmake \
    build-base \
    libx11-dev \
    pkgconf \
    sdl2-dev \
    sdl2_ttf-dev \
    sdl2_image-dev


COPY . /go/src/github.com/user/sws/
WORKDIR /go/src/github.com/user/sws/

RUN go get ./...
RUN go build ./...