package cmd

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/allenslian/vanda/infrastructure/config"
	"github.com/allenslian/vanda/infrastructure/factory/cache"
	"github.com/allenslian/vanda/infrastructure/factory/db"
	"github.com/allenslian/vanda/infrastructure/handler"
	"github.com/allenslian/vanda/infrastructure/middleware"

	//import account module
	_ "github.com/allenslian/vanda/query/account"
)

func setup(c *config.Configuration) error {
	db.InitializeDB(&db.Option{
		DefaultURI:   c.Database.DefaultURI,
		ReadonlyURI:  c.Database.ReadonlyURI,
		MaxOpenConns: c.Database.MaxOpen,
		MaxIdleConns: c.Database.MaxIdle,
	})

	defer func() {
		log.Println("Close db...")
		db.Close()
	}()

	cache.InitializeCache(&cache.Option{
		URI:        c.Cache.KVURI,
		Password:   c.Cache.KVPassword,
		CookieName: c.Cache.CookieName,
	})

	defer func() {
		log.Println("Close cache...")
		cache.Close()
	}()

	return setupWebServer(c)
}

func setupWebServer(c *config.Configuration) error {
	route := gin.Default()
	//log.SetOutput(gin.DefaultWriter)

	middleware.ConfigureTenantMiddleware(&middleware.TenantOption{
		Key:       c.Application.TenantKey,
		Mode:      c.Application.TenantMode,
		Host:      c.Network.Host,
		BlackList: c.Network.BlackList,
	})

	handler.Configure(&handler.Option{
		TenantMode: c.Application.TenantMode,
		TenantKey:  c.Application.TenantKey,
	}).Build(route)
	return route.Run(c.Network.Listen)
}
