package server

import (
	"context"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"github.com/sprakhar77/faceit/internal/config"
	"net/http"
)

func Start(ctx context.Context, cfg config.Application) {
	dependencies, err := initDependencies(cfg)
	if err != nil {
		panic(err)
	}

	router := gin.New()
	initRoutes(router, dependencies)
	run(ctx, router, cfg.Server)
}

func run(ctx context.Context, router *gin.Engine, cfg config.Server) {
	srv := &http.Server{
		Addr:    cfg.Host + ":" + cfg.Port,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("failed to start server: %s\n", err)
		}
	}()

	<-ctx.Done()

	shutdownCtx, cancelFn := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancelFn()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Fatal("Server Shutdown: ", err)
	}

	logger.Info("Server exiting")
}
