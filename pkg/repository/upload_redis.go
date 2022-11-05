package repository

import (
	"context"

	"github.com/go-redis/redis/v9"
)

var (
	ctx = context.Background()
)

type Redis struct {
	rdb *redis.Client
}

func NewUploadRedis(rdb *redis.Client) *Redis {
	return &Redis{rdb: rdb}
}

func (r *Redis) UploadConfig(cfgName string, cfgBody []byte) error {
	if err := r.rdb.Set(ctx, cfgName, cfgBody, 0).Err(); err != nil {
		return err
	}

	return nil
}
func (r *Redis) GetConfig(cfgName string) ([]byte, error) {
	var rb []byte

	rb, err := r.rdb.Get(ctx, cfgName).Bytes()
	if err != nil {
		return nil, err
	}

	return rb, nil
}
func (r *Redis) DeleteConfig(cfgName string) error {
	if err := r.rdb.Del(ctx, cfgName).Err(); err != nil {
		return err
	}

	return nil
}
