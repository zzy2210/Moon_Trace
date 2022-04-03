package web

import (
	"Moon_Trace/web/status"
	"github.com/labstack/echo"
)

type appPortRequest struct {
	Domain string // 目标域名
	Limit  int    // 起始位置，默认全部
	Offset int    // 一次性获取数量
}

type appPortResponse struct {
}

func (s *Server) AppPort(c echo.Context) error {
	return status.Success(c, "hello")
}
