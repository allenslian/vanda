package middleware

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	redis "gopkg.in/redis.v5"

	"github.com/allenslian/vanda/infrastructure/factory/cache"
)

var (
	//ErrNotFoundCache means current context has no redis client instance.
	ErrNotFoundCache = errors.New("Don't find redis client instance in the current context")
)

const (
	vandacache = "vanda_cache"
)

//CacheMiddleware is one middleware including redis client.
func CacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		client, err := cache.GetRedisCache()
		if err != nil {
			log.Printf("CacheMiddleware:GetRedisCache error:%s.\n", err.Error())
			c.Next()
			return
		}

		c.Set(vandacache, client)
		c.Next()
	}
}

// GetRedisCacheFromContext will return one redis client instance from current context.
func GetRedisCacheFromContext(c *gin.Context) (*redis.Client, bool) {
	value, exist := c.Get(vandacache)
	if !exist {
		return nil, false
	}
	cache, ok := value.(*redis.Client)
	if !ok {
		return nil, false
	}
	return cache, true
}
