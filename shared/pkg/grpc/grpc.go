package grpc

import (
	"fmt"
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func Serve(port int, register func(*grpc.Server)) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	register(server)

	return server.Serve(lis)
}
