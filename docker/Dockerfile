FROM golang:1.16.2-alpine3.13

LABEL vendor="github.com/vorstenbosch/scrapi" version="1.0.0"

ENV SCRAPI_VERSION="1.0.0"

RUN go get github.com/vorstenbosch/scrapi/scrapi

USER 9999

ENTRYPOINT [ "/go/bin/scrapi" ]