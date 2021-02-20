FROM golang:alpine

RUN apk update \
  && apk add --no-cache git \
  && go get -u github.com/cosmtrek/air \
  && go get -u github.com/google/wire/cmd/wire \
  && chmod +x ${GOPATH}/bin/air \
  && chmod +x ${GOPATH}/bin/wire

ARG ROOTDIR=/go/src/mm-server

WORKDIR ${ROOTDIR}
COPY go.mod go.sum ./
RUN go mod download
COPY . .

WORKDIR ${ROOTDIR}/cmd/http/di
RUN wire

WORKDIR ${ROOTDIR}
