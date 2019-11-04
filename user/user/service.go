package user

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/tclemos/e-gommerce/user/pkg/protocol"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Create(ctx context.Context, req *protocol.CreateUserRequest) (*protocol.UserResponse, error) {
	log.Info().Msg("User Created!")
	return &protocol.UserResponse{
		User: &protocol.User{
			Id:    "UserId",
			Name:  "UserName",
			Email: "UserEmail",
		},
	}, nil
}
