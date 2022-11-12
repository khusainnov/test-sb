package sb

import (
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/khusainnov/logging"
	"github.com/khusainnov/sbercloud/gen/pb"
	"github.com/khusainnov/sbercloud/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	logger = logging.GetLogger()
)

type Server struct {
	pb.UnimplementedUploadConfigServiceServer
	srv        *service.Service
	grpcServer *grpc.Server
}

func (s *Server) UploadConfig(ctx context.Context, req *pb.Config) (*pb.ConfigBodyResponse, error) {
	rsp, err := s.srv.UploadConfig(req)
	if err != nil {
		return nil, err
	}
	return rsp, ctx.Err()
}
func (s *Server) GetConfig(ctx context.Context, req *pb.ConfigName) (*pb.ConfigBody, error) {
	rsp, err := s.srv.GetConfig(req)
	if err != nil {
		return nil, err
	}
	return rsp, ctx.Err()
}
func (s *Server) UpdateConfig(ctx context.Context, req *pb.Config) (*pb.ConfigBodyResponse, error) {
	rsp, err := s.srv.UploadConfig(req)
	if err != nil {
		return nil, err
	}
	return rsp, ctx.Err()
}
func (s *Server) DeleteConfig(ctx context.Context, req *pb.ConfigName) (*pb.ConfigBodyResponse, error) {
	rsp, err := s.srv.DeleteConfig(req)
	if err != nil {
		return nil, err
	}
	return rsp, ctx.Err()
}

func (s *Server) RunGRPC(port string, srv *service.Service) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	s.grpcServer = grpc.NewServer()

	pb.RegisterUploadConfigServiceServer(s.grpcServer, &Server{srv: srv})
	reflection.Register(s.grpcServer)

	return s.grpcServer.Serve(lis)
}

func (s *Server) RunGateway(port string, srv *service.Service) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Fatalf("Error due start the gateway server: %+v", err)
	}

	grpcMux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err = pb.RegisterUploadConfigServiceHandlerServer(ctx, grpcMux, &Server{srv: srv}); err != nil {
		logger.Fatalf("Error due register gateway handler server: %+v", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	if err = http.Serve(lis, mux); err != nil {
		logger.Fatalf("Cannot start gateway server: %+v", err)
	}
}
