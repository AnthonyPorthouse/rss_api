FROM golang:alpine

RUN apk --no-cache --update add git openssh

RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/api

COPY ./src/api /go/src/api

RUN dep ensure --vendor-only

RUN go install

HEALTHCHECK --interval=5m --timeout=3s \
    CMD curl -f http://localhost:1234/status

CMD ["go-wrapper", "run"]