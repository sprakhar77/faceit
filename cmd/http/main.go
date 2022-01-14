package main

import (
	"context"
	logger "github.com/sirupsen/logrus"
	"github.com/sprakhar77/faceit/cmd/http/server"
	"github.com/sprakhar77/faceit/internal/config"
	"github.com/sprakhar77/faceit/internal/log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		logger.WithError(err).Fatal("failed to load config")
	}

	log.Init(cfg.Log)
	server.Start(contextWithTermSignal(), *cfg)
}

// contextWithTermSignal returns a context that will be cancelled whenever a SIGTERM is received.
func contextWithTermSignal() context.Context {
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
		<-signals
		logger.Info("Received TERM signal, stopping service...")
		cancelFunc()
	}()
	return ctx
}
