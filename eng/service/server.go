package service

import (
	"context"
	"gorm.io/gorm"

	pb "Moon_Trace/api/eng/v1"
)

type Server struct {
	pb.AppServer
	DB *gorm.DB
}

func (s *Server) HandleAppDomain(ctx context.Context, request *pb.AppDomainRequest) (*pb.AppDomainResponse, error) {
	return &pb.AppDomainResponse{
		Urls: []string{"hello"},
	}, nil
}

func (s *Server) HandleAppPort(ctx context.Context, request *pb.AppPortRequest) (*pb.AppPortResponse, error) {
	return &pb.AppPortResponse{
		Port: []int64{},
	}, nil
}

func NewServer(db *gorm.DB) *Server {
	return &Server{
		DB: db,
	}
}
