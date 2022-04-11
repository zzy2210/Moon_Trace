package web

import (
	pb "Moon_Trace/api/eng/v1"
	"Moon_Trace/web/status"
	"fmt"
	"github.com/labstack/echo"
	"google.golang.org/grpc"
	"net/http"

	"context"
)

type appPortRequest struct {
	Ip     string `json:"ip"` // 目标域名
	Limit  int    // 起始位置，默认全部
	Offset int    // 一次性获取数量
}

type appPortResponse struct {
	Ports []int
}

func (s *Server) AppPort(c echo.Context) error {
	var req appPortRequest
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
	resp, err := client.HandleAppPort(ctx, &pb.AppPortRequest{
		Ip: req.Ip,
	})
	if err != nil {
		fmt.Println("grpc err", err)
	}
	ports := []int{}
	for _, port := range resp.GetPort() {
		ports = append(ports, int(port))
	}
	return status.Success(c, &appPortResponse{
		Ports: ports,
	})
}
