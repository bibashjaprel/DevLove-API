
# syntax=docker/dockerfile:1
FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN apk add --no-cache ca-certificates && update-ca-certificates
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o devloveapi ./cmd/devloveapi/main.go

FROM scratch
WORKDIR /app
# copy only the statically linked binary and CA certs from the builder
COPY --from=builder /app/devloveapi ./devloveapi
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
EXPOSE 8080
CMD ["./devloveapi"]
