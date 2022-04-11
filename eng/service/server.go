package service

import (
	"context"
	"fmt"
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
	resp := PortScan(request.GetIp())
	ports := []int64{}
	for _, port := range resp.PortList {
		ports = append(ports, int64(port))
	}
	return &pb.AppPortResponse{
		Port: ports,
	}, nil
}

func (s *Server) HandleAppurl(ctx context.Context, request *pb.AppUrlRequest) (*pb.AppUrlResponse, error) {
	fmt.Println("url:", request.GetUrl())
	paths := FindPath(request.GetUrl())
	return &pb.AppUrlResponse{
		Urls: paths.PathList,
	}, nil
}

func NewServer(db *gorm.DB) *Server {
	return &Server{
		DB: db,
	}
}
