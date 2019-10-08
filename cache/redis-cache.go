package cache

import "github.com/go-redis/redis"

type Cache struct {
	Client *redis.Client
}

func NewRedisCache(redis *redis.Client) Cache {

	cache := Cache{
		Client: redis,
	}

	return cache
}

func (cache *Cache) Get(key string, callback Callback) string {

	result, err := cache.Client.Get(key).Result()
	if err == nil {
		if result != "" {
			return result
		}
	}

	item := NewItem()
	data := callback(&item)

	go func() {
		err = cache.Client.Set(key, data, item.TTL).Err()
		if err != nil {
			panic(err)
		}
	}()

	return data
}
