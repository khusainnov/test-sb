package sbercloud

import (
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/khusainnov/logging"
	"github.com/khusainnov/sbercloud/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	logger = logging.GetLogger()
)

type Server struct {
	pb.UnimplementedUploadConfigServiceServer
	grpcServer *grpc.Server
}

func (s *Server) RunGRPCServer(port string) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	s.grpcServer = grpc.NewServer()

	pb.RegisterUploadConfigServiceServer(s.grpcServer, &Server{})
	reflection.Register(s.grpcServer)

	return s.grpcServer.Serve(lis)
}

func (s *Server) RunGatewayServer(port string, h http.Handler) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Errorf("Cannot create listener: %s", err.Error())
		return
	}

	jsonOptions := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOptions)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err = pb.RegisterUploadConfigServiceHandlerServer(ctx, grpcMux, &Server{}); err != nil {
		logger.Fatalf("Error due registering upload handler server: %s", err.Error())
		return
	}

	err = http.Serve(lis, h)
	if err != nil {
		logger.Errorf("Error due start the gateway server: %s", err.Error())
		return
	}
}
