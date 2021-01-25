package internal

import (
	"context"
	"github.com/manifoldfinance/cabalrpc/internal/broker"
	"github.com/manifoldfinance/cabalrpc/internal/config"
	"github.com/manifoldfinance/cabalrpc/internal/log"
	"github.com/manifoldfinance/cabalrpc/internal/net"
	"go.uber.org/zap"
)

var logger *zap.Logger

type Cabalrpc interface {
	Start()
}

func NewCabalrpc(config config.Config, broker broker.Broker) (Cabalrpc, error) {
	logger = log.GetLogger(config)
	return cabalrpc{
		config:    config,
		broker:    broker,
		rpcClient: net.NewRpcClient(config),
	}, nil
}

type cabalrpc struct {
	config    config.Config
	broker    broker.Broker
	rpcClient net.RpcClient
}

func (s cabalrpc) Start() {
	defer logger.Sync()
	err := s.broker.Subscribe(s.config.TopicIncomingRpcRequests, s.onIncomingRequest)
	if err != nil {
		logger.Error("failed to subscribe to incoming request topic", zap.Error(err))
	}
}

func (s cabalrpc) onIncomingRequest(request []byte) {
	response, err := s.rpcClient.Call(request)
	if err != nil {
		s.handleError(err)
		return
	}
	s.handleError(s.broker.Publish(context.Background(), s.config.TopicOutgoingRpcResponses, response))
}

func (s cabalrpc) handleError(err error) {
	if err != nil {
		logger.Error("error occurred", zap.Error(err))
		_ = s.broker.Publish(context.Background(), s.config.TopicErrors, []byte(err.Error()))
	}
}
