package notifier

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/tclemos/e-gommerce/notifier/pkg/protocol"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) SendEmail(ctx context.Context, req *protocol.SendEmailRequest) (*protocol.SendEmailResponse, error) {
	log.Info().Interface("req", req).Msg("Sending email")
	return &protocol.SendEmailResponse{
		Id: "EmailId",
	}, nil
}
