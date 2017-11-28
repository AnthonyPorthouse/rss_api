FROM golang:alpine as builder
RUN apk --no-cache --update add git openssh
RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/api
COPY ./src/api /go/src/api

RUN dep ensure --vendor-only
RUN CGO_ENABLED=0 GOOS=linux go build -o api .

FROM alpine:latest
LABEL maintainer="Anthony Porthouse <anthony@porthou.se>"
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /go/src/api/api .
COPY --from=builder /go/src/api/migrations ./migrations

HEALTHCHECK --interval=5m --timeout=3s \
    CMD curl -f http://localhost:1234/status

CMD ["./api"]