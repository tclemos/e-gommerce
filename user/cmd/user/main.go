package main

import (
	"github.com/tclemos/e-gommerce/shared/pkg/grpc"
	"github.com/tclemos/e-gommerce/shared/pkg/log"
	"github.com/tclemos/e-gommerce/user/pkg/protocol"
	"github.com/tclemos/e-gommerce/user/user"
)

func main() {
	log.Configure()
	serve()
}

func serve() {
	if err := grpc.Serve(8080, register); err != nil {
		log.Fatal().Err(err).Msg("Failed to serve gRPC")
	}
}

func register(server *grpc.Server) {
	service := user.NewService()
	protocol.RegisterUserServiceServer(server, service)
}
