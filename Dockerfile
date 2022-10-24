# Build stage
FROM golang:1.18-alpine3.16 AS builder
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY simplebank-prod.json .

EXPOSE 9999
CMD [ "/app/main" ]