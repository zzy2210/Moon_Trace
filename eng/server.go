package eng

import (
	"Moon_Trace/eng/conf"
	"Moon_Trace/eng/model"
	"Moon_Trace/eng/service"
	"net"

	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	pb "Moon_Trace/api/eng/v1"
)

type Args struct {
	ConfigPath  string
	CertPemPath string
	CertKeyPath string
	Addr        string
}
type Server struct {
	grpcSrv *grpc.Server
	DB      *gorm.DB
	ssrv    *service.Server

	Conf *conf.Conf
	Args *Args
}

func NewSrv(DB *gorm.DB, conf *conf.Conf) *Server {
	return &Server{
		DB:   DB,
		Conf: conf,
	}
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", s.Args.Addr)
	if err != nil {
		log.Fatalf("%v", err)
	}
	grpcServer := grpc.NewServer()
	// register grpc pb
	pb.RegisterAppServer(grpcServer, s.ssrv)
	s.grpcSrv = grpcServer
	return s.grpcSrv.Serve(lis)
}

func Execute(args *Args) {
	config, err := conf.Load(args.ConfigPath)
	if err != nil {
		log.Errorf("get config error:", err)
	}
	gormDB, err := model.NewGormDB(config)
	srv := NewSrv(gormDB, config)
	srv.Args = args
	ssrv := service.NewServer(gormDB)
	srv.ssrv = ssrv
	srv.Run()
}
