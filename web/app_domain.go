package web

import (
	pb "Moon_Trace/api/eng/v1"
	"Moon_Trace/web/status"
	context "context"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"google.golang.org/grpc"
)

type appDomainRequest struct {
	Domain string `json:"domain"` // 目标域名
	Limit  int    // 起始位置，默认全部
	Offset int    // 一次性获取数量
}

type appDomainResponse struct {
	SubDomains []string `json:"subDomains"`
}

func (s *Server) AppDomain(c echo.Context) error {
	var req appDomainRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusOK, fmt.Sprintf("error:%v", err))
	}
	grpcAddr := s.Conf.Grpc.GrpcAddr[0]
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {

	}
	client := pb.NewAppClient(conn)
	ctx := context.Background()
	resp, err := client.HandleAppDomain(ctx, &pb.AppDomainRequest{
		Domain: req.Domain,
		Limit:  int64(req.Limit),
		Offset: int64(req.Offset),
	})
	if err != nil {
		fmt.Println("grpc err", err)
	}
	return status.Success(c, &appDomainResponse{
		SubDomains: resp.GetUrls(),
	})
}
