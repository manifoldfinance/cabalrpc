package config

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	BrokerType                string
	NatsUrl                   string
	KafkaUrl                  string
	HttpEnabled               bool
	HttpPort                  int
	RpcUrl                    string
	TopicIncomingRpcRequests  string
	TopicOutgoingRpcResponses string
	TopicErrors               string
	LogLevel                  LogLevelFlag
	ApmEnabled                bool
}

func NewDefaultConfig() Config {
	return Config{
		BrokerType:                BrokerNats,
		NatsUrl:                   "nats://127.0.0.1:4222",
		KafkaUrl:                  "127.0.0.1:9092",
		HttpEnabled:               false,
		HttpPort:                  8080,
		RpcUrl:                    "http://127.0.0.1:8545",
		TopicIncomingRpcRequests:  "rpc.request",
		TopicOutgoingRpcResponses: "rpc.response",
		TopicErrors:               "errors",
		LogLevel:                  NewLogLevelFlag(),
		ApmEnabled:                false,
	}
}

func (c Config) string() string {
	payload, _ := json.MarshalIndent(c, "", "\t")
	return string(payload)
}

func (c Config) ListenAddr() string {
	return fmt.Sprintf(":%d", c.HttpPort)
}
