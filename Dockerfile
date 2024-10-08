FROM golang:1.22.3-alpine AS build

RUN mkdir /opt/app
WORKDIR /opt/app

COPY go.mod go.sum ./
RUN go mod download

COPY ./*.go ./
COPY ./base_common ./base_common
COPY ./build ./build
COPY ./internal ./internal

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
RUN golangci-lint run --timeout 5m --concurrency=2

RUN go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest
RUN fieldalignment -json -fix ./...

WORKDIR /opt/app

FROM alpine:latest

RUN apk add tzdata
ENV TZ=Asia/Ho_Chi_Minh
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

CMD ["./bin"]