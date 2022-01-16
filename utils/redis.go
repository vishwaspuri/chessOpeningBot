package utils

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/vishwaspuri/ecoCodes/data"
	"time"
)

func GetCache(path string, rdb *redis.Client) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	val, err := rdb.Get(ctx, path).Result()
	if err != nil {
		return nil, err
	}

	response := data.Opening{}
	err = json.Unmarshal([]byte(val), &response)
	if err != nil {
		return nil, err
	}

	return response, err
}

func InsertCache(path string, data interface{}, rdb *redis.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	value, err := json.Marshal(data)

	err = rdb.Set(ctx, path, value, 180*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}
