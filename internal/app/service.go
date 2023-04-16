package app

import (
	pb "authserver/pkg/api"
)

type Server struct {
	repo Repository
	pb.UnimplementedAuthServerServer
}

func New(repo Repository) *Server {
	return &Server{repo: repo}
}
