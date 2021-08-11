package services

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/gommon/log"
	"github.com/rhperera/marvel-comic-api/domain"
)

type ICacheService interface {
	Connect() bool
	AddIds(ids *IdsHolder) *domain.Error
	GetIds() (string, *domain.Error)
}

type IdsHolder struct {
	Ids []int
}

func (idsHolder *IdsHolder) MarshalBinary() ([]byte, error) {
	return json.Marshal(idsHolder.Ids)
}

func (idsHolder *IdsHolder) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &idsHolder)
}

type RedisCacheService struct {
	rdb *redis.Client
}

func (cache *RedisCacheService) Connect() bool {
	cache.rdb = redis.NewClient(&redis.Options{
		Addr: "ec2-3-133-130-7.us-east-2.compute.amazonaws.com:6379",
	})
	if cache.rdb == nil {
		return false
	}
	return true
}

func (cache *RedisCacheService) AddIds(ids *IdsHolder) *domain.Error {
	err := cache.rdb.Set(context.Background(), "ids", ids, 0).Err()
	if err != nil {
		log.Error(err)
		return &domain.Error{
			Code:    3,
			Message: "Error setting in cache",
		}
	}
	return nil
}

func (cache *RedisCacheService) GetIds() (string, *domain.Error) {
	val, err := cache.rdb.Get(context.Background(), "ids").Result()

	if err != nil {
		if err.Error() == "redis: nil" {
			return "", nil
		}
		log.Error(err)
		return "", &domain.Error{
			Code:    4,
			Message: "Error getting from cache",
		}
	}
	return val, nil
}