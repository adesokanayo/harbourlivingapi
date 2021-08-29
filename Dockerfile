# Build stage
FROM golang:1.16-alpine3.13 AS builder

ENV GOOS linux
ENV CGO_ENABLED 0
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app

FROM alpine:3.14 as production
# Add certificates
RUN apk add --no-cache ca-certificates
# Copy built binary from builder
COPY --from=builder app .
# Expose port
EXPOSE 3007
# Exec built binary
CMD ./app
