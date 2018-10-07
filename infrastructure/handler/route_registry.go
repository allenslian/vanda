package handler

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	routes = make(map[string][]IRouteHandler)
)

type (
	// RouteRegistry collects all the routers under one module/controller.
	RouteRegistry func() (routeKey string, subRoutes []IRouteHandler)
)

// Registry registers to the root router.
func Registry(registry RouteRegistry) error {
	routeKey, subRoutes := registry()
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
	return nil
}

// Build indicates to generate the root router.
func Build(root gin.IRouter) {
	web := root.Group("/")
	api := root.Group("api/v1/")
	for k, v := range routes {
		for _, h := range v {
			method := h.GetMethod()
			url := fmt.Sprintf("%s/%s", k, strings.TrimLeft(h.GetURL(), "/"))
			handlers := h.GetHandlers()

			if h.GetKind() == API {
				api.Handle(method, strings.TrimLeft(url, "/"), handlers...)
			} else {
				web.Handle(method, strings.TrimLeft(url, "/"), handlers...)
			}
		}
	}
}
