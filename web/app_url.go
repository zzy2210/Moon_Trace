package web

import (
	pb "Moon_Trace/api/eng/v1"
	"Moon_Trace/web/status"
	"context"
	"fmt"
	"github.com/labstack/echo"
	"google.golang.org/grpc"
	"net/http"
)

type appUrlRequest struct {
	Domain string `json:"domain"` // 目标域名
	Limit  int    // 起始位置，默认全部
	Offset int    // 一次性获取数量
}

type appUrlResponse struct {
	Urls []string
}

func (s *Server) AppUrl(c echo.Context) error {
	var req appUrlRequest
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
	resp, err := client.HandleAppUrl(ctx, &pb.AppUrlRequest{
		Url: req.Domain,
	})
	if err != nil {
		fmt.Println("grpc err", err)
	}
	return status.Success(c, &appUrlResponse{
		Urls: resp.GetUrls(),
	})
}
