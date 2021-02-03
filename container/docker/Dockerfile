# Build in a stock Go builder container
FROM golang:1.15-alpine as builder

# TODO

RUN apk add --no-cache git make gcc musl-dev linux-headers

RUN git clone --depth 1 https://github.com/manifoldfinance/cabalrpc.git /cabal-rpc

RUN cd /cabal-rpc && make all

# Pull all binaries into a second stage deploy alpine container
FROM alpine:latest

# TODO
# PIN ALPINE VERSION
# DOCKERFILE SCHEMA
# ARG'S
# ENV


RUN apk add --no-cache ca-certificates
COPY --from=builder /go-ethereum/build/bin/* /usr/local/bin/

EXPOSE 8545 8546 30303 30303/udp
