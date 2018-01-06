FROM golang:1.9-stretch

RUN apt-get update && \
    apt-get upgrade -y
RUN apt-get install -y \
    pkg-config \
    libsdl2-dev \
    libsdl-ttf2.0-dev \
    libsdl2-image-2.0-0

COPY . /go
WORKDIR /go
#RUN go get ./...
