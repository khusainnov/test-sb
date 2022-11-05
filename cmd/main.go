package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/khusainnov/logging"
	sb "github.com/khusainnov/sbercloud"
	"github.com/khusainnov/sbercloud/driver"
	"github.com/khusainnov/sbercloud/pkg/handler"
	"github.com/khusainnov/sbercloud/pkg/repository"
	"github.com/khusainnov/sbercloud/pkg/service"
)

var (
	logger = logging.GetLogger()
	ctx    = context.Background()
)

func main() {
	if err := godotenv.Load("./config/.env"); err != nil {
		logger.Fatalf("Cannot load config, due to error: %s", err.Error())
	}

	logger.Infof("Connecting to redis on port:%s", os.Getenv("REDIS_PORT"))
	rdb, err := driver.NewRedisDB(
		driver.ConfigRedis{
			Port:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_NAME"), os.Getenv("REDIS_PORT")),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		},
		ctx,
	)
	if err != nil {
		logger.Fatalf("Cannot connect to redis, due to error: %s", err.Error())
	}

	repo := repository.NewRepository(rdb)
	services := service.NewService(repo)
	h := handler.NewHandler(services)

	s := sb.Server{}

	logger.Infof("Starting gateway server on port: %s", os.Getenv("GATE_PORT"))
	go s.RunGatewayServer(os.Getenv("GATE_PORT"), h.InitRoutes())

	logger.Infof("Starting grpc server on port: %s", os.Getenv("GRPC_PORT"))
	if err = s.RunGRPCServer(os.Getenv("GRPC_PORT")); err != nil {
		logger.Fatalf("Error due start the grpc server: %s", err.Error())
	}
}
