package cache

import (
	redis "gopkg.in/redis.v5"
)

type redisFactory struct {
	option *Option
}

func (factory *redisFactory) GetCache() func() (*redis.Client, error) {
	var client = redis.NewClient(&redis.Options{
		Addr:     factory.option.URI,
		Password: factory.option.Password,
		PoolSize: 30,
		DB:       0,
	})

	var _, err = client.Ping().Result()
	if err != nil {
		return func() (*redis.Client, error) {
			return nil, err
		}
	}
	return func() (*redis.Client, error) {
		return client, nil
	}
}

func newRedisFactory(option *Option) Factory {
	return &redisFactory{
		option: option,
	}
}

func init() {
	New = newRedisFactory
}
