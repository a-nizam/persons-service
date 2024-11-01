package app

import (
	"log/slog"

	grpcapp "github.com/a-nizam/persons-service/internal/app/grpc"
	"github.com/a-nizam/persons-service/internal/services/personlist"
	"github.com/a-nizam/persons-service/internal/storage"
)

type App struct {
	GrpcApp *grpcapp.GrpcApp
}

func New(log *slog.Logger, port int, storagePath string) *App {
	storage, err := storage.New(storagePath)
	if err != nil {
		log.Error("Failed create Storage object")
	}
	pl := personlist.New(log, storage)
	grpcApp := grpcapp.New(log, port, pl)
	return &App{GrpcApp: grpcApp}
}
