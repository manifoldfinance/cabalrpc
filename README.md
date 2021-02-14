# Cabal RPC

---
[![Go Reference](https://pkg.go.dev/badge/github.com/manifoldfinance/cabalrpc.svg)](https://pkg.go.dev/github.com/manifoldfinance/cabalrpc)
![ginkgo](https://github.com/manifoldfinance/cabalrpc/workflows/ginkgo/badge.svg?branch=master)
![go](https://github.com/manifoldfinance/cabalrpc/workflows/go/badge.svg)
---

## Overview

Kafka-based Ethereum RPC Gateway and mempool service

- [Install](#install)
- [Examples](#examples)

## Install

### `go mod`

> Requires go 1.14+

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get -u https://github.com/manifoldfinance/cabalrpc
```

## Build from source

```sh
make all
```

## Usage

| flag                  | type     | description                                                              |
| --------------------- | -------- | ------------------------------------------------------------------------ |
| --apm-enabled         |          | enable application performance monitoring using elk stack                |
| --broker-type         | string   | message broker type (nats, kafka) (default "nats")                       |
| -h,                   | --help   | help for cabalrpc                                                        |
| --http-enabled        |          | start http server for administration                                     |
| --http-port           | int      | http port (default 8080)                                                 |
| --kafka-url           | string   | kafka bootstrap server (default "127.0.0.1:9092")                        |
| --logging             | logLevel | log level (DEBUG, INFO, WARN, ERROR) (default DEBUG)                     |
| --nats-url            | string   | nats server url (default "nats://127.0.0.1:4222")                        |
| --rpc-url             | string   | ethereum rpc url (default "http://127.0.0.1:8545")                       |
| --topic-errors        | string   | topic to use for error handling (default "errors")                       |
| --topic-rpc-requests  | string   | topic to use for receiving incoming RPC requests (default "rpc.request") |
| --topic-rpc-responses | string   | topic to use for pushing RPC responses (default "rpc.response")`         |

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

> Logging, Tracing, Monitoring and Instrumenting

### APM Service

```sh
export ELASTIC_APM_SERVER_URL=https://....apm.europe-west1.gcp.cloud.es.io:443
export ELASTIC_APM_SECRET_TOKEN=secret
```

##### ELK

> see `run.sh` in the `bin/` directory

```bash
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

## Roadmap

- Enhance Tracing
- Logging Service Adapter
- Dockerized Container
- Improve TLS/SSL configuration
- ELK-stack integration

- Provide examples
- Provide better documentation
- Prometheus Support
- Grafna Support

## License

SPDX-License-Idnentifier: Apache-2.0
