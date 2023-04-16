package app

import (
	"authserver/internal/models"
	pb "authserver/pkg/api"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) RegisterRequest(ctx context.Context, req *pb.RegisterRequest) (*empty.Empty, error) {
	var data = models.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	err := s.repo.RegisterRequest(ctx, data)
	if err != nil && err != UserDoesNotExist {
		err = AlreadyExistErr
	}

	return &emptypb.Empty{}, err
}
