FROM golang:1.17.8-buster

RUN apt-get update && \
    apt-get install git && \
    apt install sudo \
    apt-get install lsof

WORKDIR /go/src/app

ADD . /go/src/app