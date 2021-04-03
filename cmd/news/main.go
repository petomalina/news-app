package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/petomalina/news-app/internal/news"
	"github.com/sethvargo/go-envconfig"
	"go.uber.org/zap"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	logger, _ := zap.NewProduction()
	logger.Sugar()
	defer logger.Sync()

	defer func() {
		done()
		if r := recover(); r != nil {
			logger.Fatal("application panic", zap.Any("panic", r))
		}
	}()

	var conf news.Config
	if err := envconfig.Process(ctx, &conf); err != nil {
		logger.Fatal("envconfig.Process", zap.Error(err))
	}

	newsServer := news.NewServer(
		logger.With(zap.String("service", "news")),
	)

	srv := http.Server{
		Handler: newsServer.Routes(),
	}

	logger.Info("starting listenAndServe", zap.String("port", conf.Port))
	err := listenAndServe(ctx, conf.Port, &srv)
	if err != nil {
		logger.Error(
			"failed to listenAndServe",
			zap.Error(err),
		)
	}

	logger.Info("successful exit")
}

func listenAndServe(ctx context.Context, port string, srv *http.Server) error {
	errCh := make(chan error, 1)
	go func() {
		<-ctx.Done()
		shutdownCtx, done := context.WithTimeout(context.Background(), 5*time.Second)
		defer done()

		err := srv.Shutdown(shutdownCtx)
		if err != nil {
			select {
			case errCh <- err:
			default:
			}
		}
	}()

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	err = srv.Serve(lis)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	select {
	case err := <-errCh:
		return fmt.Errorf("shutdown failed: %w", err)
	default:
		return nil
	}
}
