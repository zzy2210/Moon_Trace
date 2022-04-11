package service

import (
	"Moon_Trace/eng/model"
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
	sub := []string{}
	subdomains := []model.SubDomain{}
	result := s.DB.Where("domain = ?", request.GetDomain()).Find(&subdomains)
	if result.Error != nil {
		fmt.Println("err", result.Error.Error())
		return nil, nil
	}
	if result.RowsAffected == 0 {
		sub = FindSubdomain(request.GetDomain())
		for _, subDomain := range sub {
			s.DB.Create(&model.SubDomain{
				Domain:    request.GetDomain(),
				SubDomain: subDomain,
			})
		}
	} else {
		for _, subdomain := range subdomains {
			sub = append(sub, subdomain.SubDomain)
		}
	}
	return &pb.AppDomainResponse{
		Urls: sub,
	}, nil
}

func (s *Server) HandleAppPort(ctx context.Context, request *pb.AppPortRequest) (*pb.AppPortResponse, error) {
	ports := []int64{}
	dbPorts := []model.Port{}
	result := s.DB.Where("host = ?", request.GetIp()).Find(&dbPorts)
	if result.Error != nil {
		fmt.Println("err", result.Error.Error())
		return nil, nil
	}
	if result.RowsAffected == 0 {
		resp := PortScan(request.GetIp())
		for _, port := range resp.PortList {
			s.DB.Create(&model.Port{
				Port: port,
				Host: request.GetIp(),
			})
		}
		for _, port := range resp.PortList {
			ports = append(ports, int64(port))
		}
	} else {
		for _, port := range dbPorts {
			ports = append(ports, int64(port.Port))
		}
	}
	return &pb.AppPortResponse{
		Port: ports,
	}, nil
}

func (s *Server) HandleAppUrl(ctx context.Context, request *pb.AppUrlRequest) (*pb.AppUrlResponse, error) {
	dbPaths := []model.Urls{}
	pathList := []string{}
	result := s.DB.Where("host = ?", request.GetUrl()).Find(&dbPaths)
	if result.Error != nil {
		fmt.Println("err", result.Error.Error())
		return nil, nil
	}
	if result.RowsAffected == 0 {
		paths := FindPath(request.GetUrl())
		pathList = paths.PathList
		for _, path := range pathList {
			s.DB.Create(&model.Urls{
				Host: request.GetUrl(),
				Url:  path,
			})
		}
	} else {
		for _, path := range dbPaths {
			pathList = append(pathList, path.Url)
		}
	}
	return &pb.AppUrlResponse{
		Urls: pathList,
	}, nil
}

func NewServer(db *gorm.DB) *Server {
	return &Server{
		DB: db,
	}
}
