package grpcapp

import (
	"fmt"
	"log/slog"
	"net"

	personsgrpc "github.com/a-nizam/persons-service/internal/grpc/personlist"
	personListService "github.com/a-nizam/persons-service/internal/services/personlist"

	"google.golang.org/grpc"
)

type GrpcApp struct {
	Log        *slog.Logger
	GrpcServer *grpc.Server
	Port       int
}

func New(log *slog.Logger, port int, personList *personListService.PersonList) *GrpcApp {
	grpcApp := GrpcApp{
		Log:        log,
		GrpcServer: grpc.NewServer(),
		Port:       port,
	}
	personsgrpc.Register(grpcApp.GrpcServer, *personList)
	return &grpcApp
}

func (a *GrpcApp) MustRun() {
	a.Log.Info("Launching gRPC service")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.Port))
	if err != nil {
		a.Log.Error("Failed to listen port", slog.Int("port", a.Port), slog.Any("err", err))
		panic(err)
	}
	if err = a.GrpcServer.Serve(lis); err != nil {
		a.Log.Error("Failed to serve", slog.Int("port", a.Port), slog.Any("err", err))
		panic(err)
	}
}

func (a *GrpcApp) Stop() {
	a.Log.Info("Stopping service")
	a.GrpcServer.GracefulStop()
	a.Log.Info("Service is stopped")
}
