package repository

import (
	"github.com/go-redis/redis/v9"
)

type Upload interface {
	UploadConfig(cfgName string, cfgBody []byte) (string, error)
	GetConfig(cfgName string) ([]byte, error)
	DeleteConfig(cfgName string) (string, error)
}

type Repository struct {
	Upload
}

func NewRepository(rdb *redis.Client) *Repository {
	return &Repository{
		Upload: NewUploadRedis(rdb),
	}
}
