package service

import (
	"context"

	pb "Moon_Trace/api/eng/v1"
)

type Server struct {
}

func (s *Server) HandleAppDomain(ctx context.Context, request *pb.AppDomainRequest) (*pb.AppDomainResponse, error) {
	return &pb.AppDomainResponse{
		Urls: []string{},
	}, nil
}

func (s *Server) HandleAppPort(ctx context.Context, request *pb.AppPortRequest) (*pb.AppPortResponse, error) {
	return &pb.AppPortResponse{
		Port: []int64{},
	}, nil
}

func (s *Server) mustEmbedUnimplementedAppServer() {
	//TODO implement me
	panic("implement me")
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
