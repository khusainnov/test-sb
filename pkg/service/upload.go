package service

import (
	"encoding/json"
	"fmt"

	"github.com/khusainnov/sbercloud/gen/pb"
	"github.com/khusainnov/sbercloud/pkg/repository"
)

type UploadService struct {
	repo repository.Upload
}

func NewUploadService(repo repository.Upload) *UploadService {
	return &UploadService{repo: repo}
}

func (us *UploadService) UploadConfig(cfg *pb.Config) (*pb.ConfigBodyResponse, error) {
	mBody, _ := json.Marshal(cfg.Data)
	resp, err := us.repo.UploadConfig(cfg.Service, mBody)
	if err != nil {
		return nil, err
	}

	return &pb.ConfigBodyResponse{Message: resp}, nil
}
func (us *UploadService) GetConfig(cfgName *pb.ConfigName) (*pb.ConfigBody, error) {
	fmt.Printf("service: %v\n", cfgName)
	var b map[string]string
	rsp, err := us.repo.GetConfig(cfgName.Service)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(rsp, &b); err != nil {
		return nil, err
	}
	fmt.Printf("%+v", b)

	return &pb.ConfigBody{Data: b}, nil
}
func (us *UploadService) DeleteConfig(cfgName *pb.ConfigName) (*pb.ConfigBodyResponse, error) {
	rsp, err := us.repo.DeleteConfig(cfgName.Service)
	if err != nil {
		return nil, err
	}

	return &pb.ConfigBodyResponse{Message: rsp}, nil
}
