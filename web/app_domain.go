package web

import (
	"Monn_Trace_V2/web/status"
	"github.com/labstack/echo"
)

type appDomainRequest struct {
	Domain string // 目标域名
	Limit  int    // 起始位置，默认全部
	Offset int    // 一次性获取数量
}

type appDomainResponse struct {
}

func (s *Server) AppDomain(c echo.Context) error {
	return status.Success(c, "hello")
}
