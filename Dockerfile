FROM golang:1.18-alpine
MAINTAINER Jordi Riera <kender.jr@gmail.com>

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


WORKDIR /go/src/github.com/nzin/sws/
COPY . .
#RUN go get -u github.com/golang/lint/golint
RUN go get ./...
RUN go build ./...
