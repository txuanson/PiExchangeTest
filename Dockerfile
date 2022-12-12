# Stage 1: Build the binary
FROM golang:1.19.4-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main .

# # Stage 2: Build the image
FROM alpine:latest

WORKDIR /

COPY --from=builder /app/main .

ENTRYPOINT [ "./main" ]