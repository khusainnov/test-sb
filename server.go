package sb

import (
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/khusainnov/logging"
	"github.com/khusainnov/sbercloud/gen/pb"
	"github.com/khusainnov/sbercloud/pkg/endpoint"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	logger = logging.GetLogger()
)

func RunGRPC(port string, cfg *endpoint.ConfigService) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUploadConfigServiceServer(grpcServer, cfg)
	reflection.Register(grpcServer)

	return grpcServer.Serve(lis)
}

func RunGateway(port string, cfg *endpoint.ConfigService) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Fatalf("Error due start the gateway server: %+v", err)
	}

	grpcMux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err = pb.RegisterUploadConfigServiceHandlerServer(ctx, grpcMux, cfg); err != nil {
		logger.Fatalf("Error due register gateway handler server: %+v", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	if err = http.Serve(lis, mux); err != nil {
		logger.Fatalf("Cannot start gateway server: %+v", err)
	}
}
