package net

import (
	"github.com/manifoldfinance/cabalrpc/internal/broker"
	"github.com/manifoldfinance/cabalrpc/internal/config"
	"github.com/manifoldfinance/cabalrpc/internal/log"
	"github.com/gorilla/mux"
	"go.elastic.co/apm/module/apmgorilla"
	"go.elastic.co/apm/module/apmzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io/ioutil"
	"net/http"
)

var (
	logger *zap.Logger
)

type CabalrpcServer interface {
	Start()
}

func NewCabalrpcServer(config config.Config, broker broker.Broker, logLevel zapcore.Level) CabalrpcServer {
	logger = log.GetLogger(config)
	return server{
		config: config,
		broker: broker,
	}
}

type server struct {
	config config.Config
	broker broker.Broker
}

func (s server) Start() {
	logger.Info("starting cabalrpc http server")
	defer logger.Sync()
	router := mux.NewRouter()
	if s.config.ApmEnabled {
		router.Use(apmgorilla.Middleware())
	}
	router.HandleFunc("/", s.home)
	router.HandleFunc("/pub/{topic}/", s.pub)
	router.HandleFunc("/sub/{topic}/", s.sub)
	logger.Error("cannot start cabalrpc server", zap.Error(http.ListenAndServe(s.config.ListenAddr(), router)))
}

func (s server) pub(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	topic := vars["topic"]
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error("failed to read request body", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.broker.Publish(r.Context(), topic, body)
	if err != nil {
		logger.Error("failed to publish message", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func (s server) sub(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	topic := vars["topic"]
	err := s.broker.Subscribe(topic, func(message []byte) {
		logger.Info("received message", zap.String("topic", topic), zap.String("message", string(message)))
	})
	if err != nil {
		logger.Error("failed to subscribe to topic", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func (server) home(w http.ResponseWriter, r *http.Request) {
	traceContextFields := apmzap.TraceContext(r.Context())
	logger.With(traceContextFields...).Debug("handling home request")
	_, _ = w.Write([]byte("cabalrpc is up!\n"))
}
