package account

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/allenslian/vanda/infrastructure/handler"
	"github.com/allenslian/vanda/infrastructure/middleware"
)

func getTenants(c *gin.Context) {
	db, ok := middleware.GetReadonlyDB(c)
	if ok {
		c.AbortWithError(http.StatusInternalServerError, middleware.ErrNotFoundReadonlyDB)
	}

	db.Query(`select id, name from account.tenant;`)
}

func init() {
	handler.Registry(func() (routeKey string, subRoutes []handler.IRouteHandler) {
		handlers := []handler.IRouteHandler{
			handler.NewRoute(handler.API, "tenant", "GET", getTenants,
				middleware.ReadonlyDBMiddleware()),
		}
		return "account", handlers
	})
}
