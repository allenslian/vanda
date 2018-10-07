package cmd

import (
	"github.com/gin-gonic/gin"

	"github.com/allenslian/vanda/infrastructure/config"
	"github.com/allenslian/vanda/infrastructure/factory/cache"
	"github.com/allenslian/vanda/infrastructure/factory/db"
	"github.com/allenslian/vanda/infrastructure/handler"
)

func setup(c *config.Configuration) error {
	db.InitializeDB(&db.Option{
		DefaultURI:   c.Database.DefaultURI,
		ReadonlyURI:  c.Database.ReadonlyURI,
		MaxOpenConns: c.Database.MaxOpen,
		MaxIdleConns: c.Database.MaxIdle,
	})

	cache.InitializeCache(&cache.Option{
		URI:        c.Cache.KVURI,
		Password:   c.Cache.KVPassword,
		CookieName: c.Cache.CookieName,
	})

	return setupWebServer(c)
}

func setupWebServer(c *config.Configuration) error {
	route := gin.Default()
	handler.Build(route)
	return route.Run(c.Network.Listen)
}
