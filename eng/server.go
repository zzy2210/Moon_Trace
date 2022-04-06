package eng

import (
	"Moon_Trace/eng/conf"
	"Moon_Trace/eng/model"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type Args struct {
	ConfigPath string
}
type Server struct {
	DB *gorm.DB
}

func NewSrv(DB *gorm.DB) *Server {
	return &Server{
		DB: DB,
	}
}

func (s *Server) Run() error {
	// 这里创建grpc server

	return nil
}
func Execute(args *Args) {
	config, err := conf.Load(args.ConfigPath)
	if err != nil {
		log.Errorf("get config error:", err)
	}
	gormDB, err := model.NewGormDB(config)
	srv := NewSrv(gormDB)

	srv.Run()
}
