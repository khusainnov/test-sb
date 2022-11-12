package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/khusainnov/logging"
	"github.com/khusainnov/sbercloud"
	"github.com/khusainnov/sbercloud/driver"
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
			// if you start with docker-compose, change from localhost to Getenv
			Port:     fmt.Sprintf("%s:%s" /*os.Getenv("REDIS_NAME")*/, "localhost", os.Getenv("REDIS_PORT")),
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

	s := sb.Server{}

	logger.Infof("Starting Gateway server on port:%s", os.Getenv("GATE_PORT"))
	go s.RunGateway(os.Getenv("GATE_PORT"), services)

	logger.Infof("Starting gRPC server on port:%s", os.Getenv("GRPC_PORT"))
	if err = s.RunGRPC(os.Getenv("GRPC_PORT"), services); err != nil {
		logger.Fatalf("Cannot start the grpc server: %+v", err)
	}
}
