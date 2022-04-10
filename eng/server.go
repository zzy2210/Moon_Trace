package eng

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"

	"Moon_Trace/cert"
	"Moon_Trace/eng/conf"
	"Moon_Trace/eng/model"
	"Moon_Trace/eng/service"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	grpcSrv *http.Server
	DB      *gorm.DB

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
	conn, err := net.Listen("tcp", s.Args.Addr)
	if err != nil {
		log.Fatalf("%v", err)
	}
	tlsConf := cert.GetTLSConfig(s.Args.CertPemPath, s.Args.CertKeyPath)
	return s.grpcSrv.Serve(tls.NewListener(conn, tlsConf))
}

func Execute(args *Args) {
	config, err := conf.Load(args.ConfigPath)
	if err != nil {
		log.Errorf("get config error:", err)
	}
	gormDB, err := model.NewGormDB(config)
	srv := NewSrv(gormDB, config)
	srv.Args = args
	tlsConf := cert.GetTLSConfig(args.CertPemPath, args.CertKeyPath)
	grpcSrv, err := newGrpc(tlsConf, args)
	if err != nil {
		log.Errorf("")
	}
	srv.grpcSrv = grpcSrv
	srv.Run()
}

func newGrpc(tlsConfig *tls.Config, args *Args) (*http.Server, error) {
	var opts []grpc.ServerOption
	// grpc server
	creds, err := credentials.NewServerTLSFromFile(args.CertPemPath, args.CertKeyPath)
	if err != nil {
		fmt.Println(1)
		log.Printf("Failed to create server TLS credentials %v", err)
		return nil, err
	}
	opts = append(opts, grpc.Creds(creds))
	grpcServer := grpc.NewServer(opts...)
	// register grpc pb
	pb.RegisterAppServer(grpcServer, service.NewServer())
	// gw server
	ctx := context.Background()
	dcreds, err := credentials.NewClientTLSFromFile(args.CertPemPath, "test")
	if err != nil {
		fmt.Println(2)
		log.Printf("Failed to create client TLS credentials %v", err)
		return nil, err
	}
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}
	gwmux := runtime.NewServeMux()
	// register grpc-gateway pb
	if err := pb.RegisterAppHandlerFromEndpoint(ctx, gwmux, args.Addr, dopts); err != nil {
		fmt.Println(3)
		log.Printf("Failed to register gw server: %v\n", err)
	}
	// http服务
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	return &http.Server{
		Addr:      args.Addr,
		TLSConfig: tlsConfig,
	}, nil
}
