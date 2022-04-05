package eng

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"

	pb "Moon_Trace/api/eng/v1"
)

type Server struct {
}

func NewSrv() *Server {
	return &Server{}
}

func (s *Server) Run() error {
	ctx := context.Background()
	gwmux := runtime.ServeMux{}
	err := pb.RegisterAppDomainServer(ctx, gwmux, "")
}
func Execute() {
	srv := NewSrv()
	srv.Run()
}
