package handler

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	routes     = make(map[string][]IRouteHandler)
	registries = []RouteRegistry{}
	locker     = new(sync.Mutex)
)

type (
	// RouteRegistry collects all the routers under one module/controller.
	RouteRegistry func() (routeKey string, subRoutes []IRouteHandler)

	//Option describes some configuration items for building routes.
	Option struct {
		TenantMode string
		TenantKey  string
	}

	//RouteBuilder builds router table for the application
	RouteBuilder interface {
		Build(root gin.IRouter)
	}

	routeBuilder struct {
		option *Option
	}
)

// Registry registers to the root router.
func Registry(registry RouteRegistry) {
	locker.Lock()
	registries = append(registries, registry)
	locker.Unlock()
}

// Configure sets some configuration items for building routes.
func Configure(option *Option) RouteBuilder {
	return &routeBuilder{option: option}
}

// Build indicates to generate the root router.
func (r *routeBuilder) Build(root gin.IRouter) {
	web := root.Group("/")
	api := root.Group("api/v1/")

	buildRegistries()
	buildRoutes(r, web, api)

	// clear registry
	registries = nil
}

func buildRegistries() {
	log.Printf("build registries(%d).\n", len(registries))
	for _, v := range registries {
		routeKey, subRoutes := v()

		// if routeKey == "" {
		// 	return ErrEmptyRouteKeyNotAllowed
		// }

		key := strings.Trim(routeKey, "/")
		if _, ok := routes[key]; ok {
			//merge two subRoutes into one key.
			routes[key] = append(routes[key], subRoutes...)
		} else {
			routes[key] = subRoutes
		}
	}
}

func buildRoutes(b *routeBuilder, web, api *gin.RouterGroup) {
	log.Printf("build routes(%d).\n", len(routes))
	for k, v := range routes {
		for _, h := range v {
			method := h.GetMethod()
			handlers := h.GetHandlers()
			url := fmt.Sprintf("%s/%s", k, strings.TrimLeft(h.GetURL(), "/"))

			if h.GetKind() == API {
				if strings.ToLower(b.option.TenantMode) == "url" {
					// add one route key to api root route.
					api.Handle(method,
						fmt.Sprintf("%s/%s", ":"+b.option.TenantKey, strings.TrimLeft(url, "/")),
						handlers...)
				} else {
					api.Handle(method, strings.TrimLeft(url, "/"), handlers...)
				}
			} else {
				web.Handle(method, strings.TrimLeft(url, "/"), handlers...)
			}
		}
	}
}
