package service

import (
	"github.com/khusainnov/sbercloud/gen/pb"
	"github.com/khusainnov/sbercloud/pkg/repository"
)

type Upload interface {
	UploadConfig(config *pb.Config) (*pb.ConfigBodyResponse, error)
	GetConfig(cfgName *pb.ConfigName) (*pb.ConfigBody, error)
	DeleteConfig(cfgName *pb.ConfigName) (*pb.ConfigBodyResponse, error)
}

type Service struct {
	Upload
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Upload: NewUploadService(repo),
	}
}
