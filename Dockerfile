FROM golang:1.19 as builder

WORKDIR /go/src/greencity

ENV GO111MODULE=on

COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o producer ./cmd/producer/main.go
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o consumer ./cmd/consumer/main.go

FROM debian:buster-slim

COPY --from=builder /go/src/greencity/producer /usr/bin
COPY --from=builder /go/src/greencity/consumer /usr/bin

ENTRYPOINT ["producer"]