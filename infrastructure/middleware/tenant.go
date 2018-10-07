package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/allenslian/vanda/infrastructure/config"
)

//TenantMiddleware checks current tenant information.
func TenantMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, ok := GetReadonlyDB(c)
		if !ok {
			c.AbortWithError(http.StatusInternalServerError, ErrNotFoundReadonlyDB)
		}
		db.QueryRow(`SELECT id, name FROM account.tenant WHERE name=$1;`, "")
	}
}

func parseTenantName(c *gin.Context) string {
	var cfg = config.Get()
	if cfg == nil {
		return ""
	}
	if cfg.Application.TenantMode == "domain" {
		return strings.Replace(c.Request.Host, cfg.Network.Host, "", -1)
	} else if cfg.Application.TenantMode == "header" {

	} else {
		// url
	}
	return ""
}
