package service

import (
	"context"

	pb "Moon_Trace/api/eng/v1"
)

type Server struct {
}

func (s *Server) Echo(ctx context.Context, request *pb.AppDomainRequest) (*pb.AppDomainResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) mustEmbedUnimplementedAppDomainServer() {
	//TODO implement me
	panic("implement me")
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) AppDomain(ctx context.Context, r *pb.AppDomainRequest) (*pb.AppDomainResponse, error) {

	return &pb.AppDomainResponse{
		Urls: []string{},
	}, nil
}
