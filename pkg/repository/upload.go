package repository

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
)

var (
	ctx = context.Background()
)

type UploadRedis struct {
	rdb *redis.Client
}

func NewUploadRedis(rdb *redis.Client) *UploadRedis {
	return &UploadRedis{rdb: rdb}
}

func (ur *UploadRedis) UploadConfig(cfgName string, cfgBody []byte) (string, error) {
	if err := ur.rdb.Set(ctx, cfgName, cfgBody, 0).Err(); err != nil {
		return "", err
	}

	return "CREATED", nil
}
func (ur *UploadRedis) GetConfig(cfgName string) ([]byte, error) {
	fmt.Printf("repo: %v\n", cfgName)
	var rsp []byte

	rsp, err := ur.rdb.Get(ctx, cfgName).Bytes()
	if err != nil {
		return nil, err
	}

	return rsp, nil
}
func (ur *UploadRedis) DeleteConfig(cfgName string) (string, error) {
	if err := ur.rdb.Del(ctx, cfgName).Err(); err != nil {
		return "", err
	}

	return "DELETED", nil
}
