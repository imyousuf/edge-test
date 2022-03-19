FROM golang:1.17.8-alpine3.15 AS build-env

RUN apk update && apk add make bash

RUN mkdir -p /go-code
WORKDIR /go-code/

ADD go.mod .
ADD go.sum .
RUN go mod download

ADD main.go .

RUN go build -mod=readonly

FROM alpine:3.15
WORKDIR /
COPY --from=build-env /go-code/go-code /go-code
CMD [ "/go-code" ]
