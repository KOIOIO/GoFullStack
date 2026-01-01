FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .

# Switch to api directory to build
WORKDIR /app/api
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod tidy
RUN go build -o /app/api_server loginregister.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/api_server .
COPY --from=builder /app/api/etc/loginregister.yaml ./etc/

# Replace RPC endpoint
RUN sed -i 's/127.0.0.1:9100/go-rpc:9100/g' etc/loginregister.yaml

CMD ["./api_server"]
