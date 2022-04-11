package web

import (
	config "Moon_Trace/web/conf"
	"fmt"
	"github.com/labstack/echo"
	"log"
)

type Args struct {
	CertPemPath string
	ConfPath    string
}
type Server struct {
	e    *echo.Echo
	Conf *config.Conf

	Args *Args
	// grpcSrvs[]
}

func NewServer(e *echo.Echo, conf *config.Conf) *Server {
	return &Server{
		e:    e,
		Conf: conf,
	}
}

func (s *Server) Run() error {
	s.Register()
	port := fmt.Sprintf(":%v", s.Conf.Web.Port)
	return s.e.Start(port)
}

func Execute(args *Args) error {
	e := echo.New()
	conf, err := config.Load(args.ConfPath)
	if err != nil {
		log.Fatalln(err)
	}
	srv := NewServer(e, conf)
	srv.Args = args
	return srv.Run()
}

func (s *Server) Register() {
	s.e.POST("/app/domain", s.AppDomain)
	s.e.POST("/app/url", s.AppUrl)
	s.e.POST("/app/port", s.AppPort)
}
