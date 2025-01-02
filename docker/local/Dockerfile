FROM golang:1.18-alpine AS builder

WORKDIR /lyfebloc-edge

ADD go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o lyfebloc-edge main.go

FROM alpine:latest AS runner

RUN apk --no-cache add ca-certificates jq

WORKDIR /lyfebloc-edge

COPY --from=builder /lyfebloc-edge/lyfebloc-edge ./
COPY ./docker/local/lyfebloc-edge.sh ./

# Expose json-rpc, libp2p and grpc ports
EXPOSE 8545 9632 1478 5001

ENTRYPOINT ["./lyfebloc-edge.sh"]
