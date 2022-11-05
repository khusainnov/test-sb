package service

import (
	"github.com/khusainnov/sbercloud/internal/entity"
	"github.com/khusainnov/sbercloud/pkg/repository"
)

type Upload interface {
	UploadConfig(cfg entity.UploadData) error
	GetConfig(cfgName string) (map[string]string, error)
	DeleteConfig(cfgName string) error
}

type Service struct {
	Upload
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Upload: NewUploadService(repo),
	}
}
