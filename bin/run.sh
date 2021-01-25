#!/bin/bash
./cabalrpc \
--topic-rpc-requests=ethereum.rpc.requests \
--topic-rpc-responses=ethereum.rpc.responses \
--topic-errors=cabalrpc.errors \
--logging=INFO \
--http-enabled --http-port=9000 \
--rpc-url=http://127.0.0.1:8545 \
--nats-url=nats://127.0.0.1:4222 --broker-type=nats \
--apm-enabled
