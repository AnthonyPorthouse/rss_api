FROM golang:alpine as builder

LABEL maintainer="Anthony Porthouse <anthony@porthou.se>"

RUN apk --no-cache --update add git openssh
RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/app
COPY ./src/app /go/src/app

RUN dep ensure --vendor-only
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM node:alpine as js-builder

COPY ./src/app /app

RUN yarn && yarn --production

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /go/src/app/app .
COPY --from=builder /go/src/app/resources/templates /app/resources/templates
COPY --from=js-builder /app/public/css /app/public/css
COPY --from=js-builder /app/public/js /app/public/js

HEALTHCHECK --interval=5m --timeout=3s \
    CMD curl -f http://localhost:1234/status

CMD ["./app"]