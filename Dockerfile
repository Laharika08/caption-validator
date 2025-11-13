
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod init solution && go mod tidy && go build -o /app/caption-validator ./cmd

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/caption-validator .
ENTRYPOINT ["./caption-validator"]
