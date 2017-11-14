FROM golang:alpine

RUN apk --no-cache --update add git openssh

RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/app

COPY ./src/app /go/src/app

RUN dep ensure

RUN go install

HEALTHCHECK --interval=5m --timeout=3s \
    CMD curl -f http://localhost:1234/status

CMD ["go-wrapper", "run"]