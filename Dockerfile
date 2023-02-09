FROM golang:1.18-alpine AS builder

RUN apk add --no-cache git

WORKDIR /go-app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -x -ldflags "-s -w" -o /go-app/app

FROM alpine:3.16
RUN apk --no-cache add tzdata

WORKDIR /go-app

COPY --from=builder /go-app/app .

ENV KORDIS_USERNAME="" \
   KORDIS_PASSWORD="" \
   SCHEDULER_CRON="0 * * * *" \
   PLANNING_DAYS_SYNC=15 \
   CALENDAR_ID="" \
   TZ="Europe/Paris" \
   MODE="sync"

#sync/scheduler

VOLUME /go-app/auth

ENTRYPOINT ["/go-app/app"]
