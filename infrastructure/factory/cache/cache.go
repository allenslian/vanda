package cache

import redis "gopkg.in/redis.v5"

type (
	//Option describes some cache settings.
	Option struct {
		URI        string
		Password   string
		CookieName string
	}

	//Factory creates one cache instance.
	Factory interface {
		GetCache() func() (*redis.Client, error)
	}
)

var (
	//New creates one cache factory.
	New func(option *Option) Factory
)

var (
	factory Factory
)

//InitializeCache sets some redis settings.
func InitializeCache(option *Option) {
	factory = New(option)
}
