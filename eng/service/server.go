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
	// result := s.DB.Model(&model.SubDomain{})
	sub := FindSubdomain(request.GetDomain())
	return &pb.AppDomainResponse{
		Urls: sub,
	}, nil
}

func (s *Server) HandleAppPort(ctx context.Context, request *pb.AppPortRequest) (*pb.AppPortResponse, error) {
	return &pb.AppPortResponse{
		Port: []int64{},
	}, nil
}

func (s *Server) HandleAppurl(ctx context.Context, request *pb.AppUrlRequest) (*pb.AppUrlResponse, error) {
	return nil, nil
}

func NewServer(db *gorm.DB) *Server {
	return &Server{
		DB: db,
	}
}
