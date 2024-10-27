package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/a-nizam/persons-service/internal/app"
	"github.com/a-nizam/persons-service/internal/config"
)

func main() {
	cfg := config.MustLoad("../config/config.yaml")
	log := setupLogger()
	application := app.New(log, cfg.Grpc.Port, cfg.StoragePath)
	go application.GrpcApp.MustRun()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	application.GrpcApp.Stop()
}

func setupLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
}
