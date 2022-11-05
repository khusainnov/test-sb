package service

import (
	"encoding/json"

	"github.com/khusainnov/sbercloud/internal/entity"
	"github.com/khusainnov/sbercloud/pkg/repository"
)

type UploadService struct {
	repo repository.Upload
}

func NewUploadService(repo repository.Upload) *UploadService {
	return &UploadService{repo: repo}
}

func (us *UploadService) UploadConfig(cfg entity.UploadData) error {
	body, _ := json.Marshal(cfg.Data)

	if err := us.repo.UploadConfig(cfg.Service, body); err != nil {
		return err
	}

	return nil
}

func (us *UploadService) GetConfig(cfgName string) (map[string]string, error) {
	var rb map[string]string

	body, err := us.repo.GetConfig(cfgName)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &rb)
	if err != nil {
		return nil, err
	}

	return rb, nil
}
func (us *UploadService) DeleteConfig(cfgName string) error {
	if err := us.repo.DeleteConfig(cfgName); err != nil {
		return err
	}

	return nil
}
