package account

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/allenslian/vanda/infrastructure/handler"
	"github.com/allenslian/vanda/infrastructure/middleware"
)

func getTenants(c *gin.Context) {
	// db, ok := middleware.GetReadonlyDB(c)
	// if ok {
	// 	c.AbortWithError(http.StatusInternalServerError, middleware.ErrNotFoundReadonlyDB)
	// }
	log.Println("getTenants")
	identity, _ := middleware.GetTenantIdentityFromContext(c)
	c.JSON(http.StatusOK, identity)
}

func init() {
	handler.Registry(func() (routeKey string, subRoutes []handler.IRouteHandler) {
		handlers := []handler.IRouteHandler{
			handler.NewRoute(handler.API, "tenants", "GET", getTenants,
				middleware.ReadonlyDBMiddleware(),
				middleware.CacheMiddleware(),
				middleware.TenantMiddleware()),
		}
		return "account", handlers
	})
}
