# Cabalrpc

Cabalrpc provides an Ethereum RPC gateway over Kafka and other messaging systems (e.g. NATS)

- [Install](#install)
- [Examples](#examples)

## Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get -u https://github.com/manifoldfinance/cabalrpc
```

## Build from source

```sh
make all
```

## Usage

````
Usage:
  cabalrpc [flags]

Flags:
      --apm-enabled                  enable application performance monitoring using elk stack
      --broker-type string           message broker type (nats, kafka) (default "nats")
  -h, --help                         help for cabalrpc
      --http-enabled                 start http server for administration
      --http-port int                http port (default 8080)
      --kafka-url string             kafka bootstrap server (default "127.0.0.1:9092")
      --logging logLevel             log level (DEBUG, INFO, WARN, ERROR) (default DEBUG)
      --nats-url string              nats server url (default "nats://127.0.0.1:4222")
      --rpc-url string               ethereum rpc url (default "http://127.0.0.1:8545")
      --topic-errors string          topic to use for error handling (default "errors")
      --topic-rpc-requests string    topic to use for receiving incoming RPC requests (default "rpc.request")
      --topic-rpc-responses string   topic to use for pushing RPC responses (default "rpc.response")```
````

## Examples

Run NATS server.

```sh
docker-compose -f build/package/nats/docker-compose.yml up -d
```

Run Cabal-RPC

```sh
cabalrpc \
--topic-rpc-requests=ethereum.rpc.requests \
--topic-rpc-responses=ethereum.rpc.responses \
--topic-errors=cabalrpc.errors \
--logging=INFO \
--http-enabled --http-port=9000 \
--rpc-url=http://127.0.0.1:8545 \
--nats-url=nats://127.0.0.1:4222 --broker-type=nats \
--apm-enabled
```

## Application performance monitoring

Configure APM server.

```sh
export ELASTIC_APM_SERVER_URL=https://....apm.europe-west1.gcp.cloud.es.io:443
export ELASTIC_APM_SECRET_TOKEN=secret
```

```sh
ELASTIC_APM_SERVICE_NAME=cabalrpc cabalrpc \
--topic-rpc-requests=ethereum.rpc.requests \
--topic-rpc-responses=ethereum.rpc.responses \
--topic-errors=cabalrpc.errors \
--logging=INFO \
--http-enabled --http-port=9000 \
--rpc-url=http://127.0.0.1:8545 \
--nats-url=nats://127.0.0.1:4222 --broker-type=nats \
--apm-enabled
```

## License

Apache-2.0
