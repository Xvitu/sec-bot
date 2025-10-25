# Etapa 1: build
FROM golang:1.25.3-alpine AS builder

WORKDIR /app

# todo - add go.sum too
COPY go.mod ./
RUN go mod download

COPY .. .

RUN go build -o server .

FROM alpine:latest
EXPOSE 8080

CMD ["./server"]


WORKDIR /root/

COPY --from=builder /app/server .
