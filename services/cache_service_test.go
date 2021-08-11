package services

import (
	"github.com/rhperera/marvel-comic-api/config"
	"testing"
)

func TestRedisCacheService_AddIds(t *testing.T) {
	config.InitForTests()
	rdb := &RedisCacheService{}
	rdb.Connect()
	ids := IdsHolder{Ids: []int{1,3,4,5}}
	rdb.AddIds(&ids)
	data, err := rdb.GetIds()
	if err != nil {
		t.Error(err)
	}
	if data != "[1,3,4,5]" {
		t.Error(data)
	}
}