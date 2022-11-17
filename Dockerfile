FROM golang:1.19-alpine AS builder

RUN apk add --no-cache git

WORKDIR /go-app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -x -ldflags "-s -w" -o /go-app/app

FROM alpine:3.16


WORKDIR /go-app

COPY --from=builder /go-app/app .

ENV username=""\
    password=""

ENTRYPOINT ["/go-app/app"]
