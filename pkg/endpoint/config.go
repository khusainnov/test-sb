package endpoint

import (
	"context"

	"github.com/khusainnov/sbercloud/gen/pb"
	"github.com/khusainnov/sbercloud/pkg/service"
)

type ConfigService struct {
	pb.UnimplementedUploadConfigServiceServer
	srv *service.Service
}

func NewConfigService(srv *service.Service) *ConfigService {
	return &ConfigService{srv: srv}
}

func (cs *ConfigService) UploadConfig(ctx context.Context, req *pb.Config) (*pb.ConfigBodyResponse, error) {
	rsp, err := cs.srv.UploadConfig(req)
	if err != nil {
		return nil, err
	}
	return rsp, ctx.Err()
}
func (cs *ConfigService) GetConfig(ctx context.Context, req *pb.ConfigName) (*pb.ConfigBody, error) {
	rsp, err := cs.srv.GetConfig(req)
	if err != nil {
		return nil, err
	}
	return rsp, ctx.Err()
}
func (cs *ConfigService) UpdateConfig(ctx context.Context, req *pb.Config) (*pb.ConfigBodyResponse, error) {
	rsp, err := cs.srv.UploadConfig(req)
	if err != nil {
		return nil, err
	}
	return rsp, ctx.Err()
}
func (cs *ConfigService) DeleteConfig(ctx context.Context, req *pb.ConfigName) (*pb.ConfigBodyResponse, error) {
	rsp, err := cs.srv.DeleteConfig(req)
	if err != nil {
		return nil, err
	}
	return rsp, ctx.Err()
}
