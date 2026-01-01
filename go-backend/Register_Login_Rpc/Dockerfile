FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .

RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod tidy
RUN go build -o /app/rpc_server loginregister.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/rpc_server .

# Create config file dynamically to avoid Etcd dependency and fix Mongo host
# AuthMechanism is hardcoded in the Go code, so we don't need it in the config file (and adding it would cause a config load error)
RUN mkdir -p etc && \
    printf "Name: loginregister.rpc\nListenOn: 0.0.0.0:9100\n\nMongo:\n  Hosts:\n    - mongo:27017\n  Database: login_register\n  User: wwy\n  Password: \"123456\"\n  AuthSource: admin\n" > etc/loginregister.yaml

CMD ["./rpc_server"]
