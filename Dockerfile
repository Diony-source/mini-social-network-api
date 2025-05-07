FROM golang:1.23.4-alpine3.18 AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/app .

CMD ["./app"]
