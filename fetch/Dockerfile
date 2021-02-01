FROM golang:alpine AS builder

RUN apk add --update --no-cache ca-certificates git
ENV GO111MODULE=on

RUN  mkdir  /api
WORKDIR /api/
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /api/fetch

ENTRYPOINT ["/api/fetch"]