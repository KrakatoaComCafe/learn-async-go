FROM golang:1.24.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/app

FROM alpine:latest

RUN apk add --no-cache curl && \
    curl -sSL https://github.com/jwilder/dockerize/releases/download/v0.6.1/dockerize-linux-amd64-v0.6.1.tar.gz \
    | tar -C /usr/local/bin -xzv

WORKDIR /app
COPY --from=builder /app/app .

EXPOSE 8080
CMD ["dockerize", "-wait", "tcp://kafka:9092", "-timeout", "30s", "./app"]