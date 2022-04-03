package web

import "github.com/labstack/echo"

type Server struct {
	e *echo.Echo
}

func NewServer(e *echo.Echo) *Server {
	return &Server{
		e: e,
	}
}

func Start() error {
	e := echo.New()
	srv := NewServer(e)
	return srv.Run()
}
func (s *Server) Run() error {
	s.Register()

	return s.e.Start(":1323")
}

func (s *Server) Register() {
	s.e.GET("/app/domain", s.AppDomain)
}
