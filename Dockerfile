FROM golang:1.17.8-buster

RUN apt-get update && apt-get install git

WORKDIR /go/src/app

ADD . /go/src/app