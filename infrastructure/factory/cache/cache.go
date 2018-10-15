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
		Close() error
	}
)

var (
	//New creates one cache factory.
	New func(option *Option) Factory

	factory Factory
)

//InitializeCache sets some redis settings.
func InitializeCache(option *Option) {
	factory = New(option)
}

//GetRedisCache get one instance of redis client.
func GetRedisCache() (*redis.Client, error) {
	if factory == nil {
		return nil, errMissingOption
	}
	return factory.GetCache()()
}

//Close will close the client.
func Close() error {
	if factory == nil {
		return errMissingOption
	}
	return factory.Close()
}
