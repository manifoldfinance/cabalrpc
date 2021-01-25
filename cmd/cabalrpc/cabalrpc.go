package main

import (
	"github.com/manifoldfinance/cabalrpc/internal"
	"github.com/manifoldfinance/cabalrpc/internal/config"
	"github.com/manifoldfinance/cabalrpc/internal/log"
	"github.com/manifoldfinance/cabalrpc/internal/net"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/cobra"
	"go.elastic.co/apm/module/apmot"
	"go.uber.org/zap"
	"os"
)

func main() {
	cabalrpcConfig := config.NewDefaultConfig()

	cmd := &cobra.Command{
		Use:   "cabalrpc",
		Short: "cabalrpc provides an Ethereum RPC gateway over Kafka.",
		RunE:  run(&cabalrpcConfig),
	}
	cmd.PersistentFlags().StringVar(&cabalrpcConfig.BrokerType, "broker-type", cabalrpcConfig.BrokerType, "message broker type (nats, kafka)")
	cmd.PersistentFlags().StringVar(&cabalrpcConfig.KafkaUrl, "kafka-url", cabalrpcConfig.KafkaUrl, "kafka bootstrap server")
	cmd.PersistentFlags().StringVar(&cabalrpcConfig.NatsUrl, "nats-url", cabalrpcConfig.NatsUrl, "nats server url")
	cmd.PersistentFlags().BoolVar(&cabalrpcConfig.HttpEnabled, "http-enabled", cabalrpcConfig.HttpEnabled, "start http server for administration")
	cmd.PersistentFlags().IntVar(&cabalrpcConfig.HttpPort, "http-port", cabalrpcConfig.HttpPort, "http port")
	cmd.PersistentFlags().StringVar(&cabalrpcConfig.RpcUrl, "rpc-url", cabalrpcConfig.RpcUrl, "ethereum rpc url")
	cmd.PersistentFlags().StringVar(&cabalrpcConfig.TopicIncomingRpcRequests, "topic-rpc-requests", cabalrpcConfig.TopicIncomingRpcRequests, "topic to use for receiving incoming RPC requests")
	cmd.PersistentFlags().StringVar(&cabalrpcConfig.TopicOutgoingRpcResponses, "topic-rpc-responses", cabalrpcConfig.TopicOutgoingRpcResponses, "topic to use for pushing RPC responses")
	cmd.PersistentFlags().StringVar(&cabalrpcConfig.TopicErrors, "topic-errors", cabalrpcConfig.TopicErrors, "topic to use for error handling")
	cmd.PersistentFlags().Var(&cabalrpcConfig.LogLevel, "logging", "log level (DEBUG, INFO, WARN, ERROR)")
	cmd.PersistentFlags().BoolVar(&cabalrpcConfig.ApmEnabled, "apm-enabled", cabalrpcConfig.ApmEnabled, "enable application performance monitoring using elk stack")

	err := cmd.Execute()
	logger := log.GetLogger(cabalrpcConfig)
	defer logger.Sync()
	if err != nil {
		logger.Error("Failed to execute", zap.Error(err))
		os.Exit(1)
	}
}

func run(cfg *config.Config) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if cfg.ApmEnabled {
			opentracing.SetGlobalTracer(apmot.New())
		}
		broker, err := internal.NewBroker(*cfg)
		if err != nil {
			return err
		}
		cabalrpc, err := internal.NewCabalrpc(*cfg, broker)
		if err != nil {
			return err
		}
		go cabalrpc.Start()
		if cfg.HttpEnabled {
			net.NewCabalrpcServer(*cfg, broker, cfg.LogLevel.ZapLevel).Start()
		}

		return nil
	}
}
