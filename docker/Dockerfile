FROM golang:1.16.2-alpine3.13

LABEL vendor="github.com/vorstenbosch/scrapi" version="1.0.0"

RUN go get github.com/vorstenbosch/scrapi/scrapi

ENTRYPOINT [ "${GOPATH}/bin/scrapi" ]