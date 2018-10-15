package middleware

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	redis "gopkg.in/redis.v5"

	"github.com/allenslian/vanda/infrastructure/utils"
)

const (
	tenantContextKey = "tenant.context"
)

var (
	errMissingOption          = errors.New("Missing tenant option object")
	errMissingTenantKeyValue  = errors.New("The application's tenant_key is missing")
	errNotImplementTenantMode = errors.New("Not implement the tenant mode")
	errInvalidTenantAccount   = errors.New("Invalid tenant account")
)

var (
	_option *TenantOption
)

//TenantIdentity describes one basic tenant model.
type TenantIdentity struct {
	ID      string `json:"id"`
	Account string `json:"account"`
}

//TenantOption describes some configruation items.
type TenantOption struct {
	Mode      string
	Key       string
	Host      string
	BlackList []string
}

//ConfigureTenantMiddleware sets tenant middleware's option.
func ConfigureTenantMiddleware(option *TenantOption) {
	_option = option
}

//TenantMiddleware checks current tenant information.
func TenantMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		account, err := parseTenantAccount(c, _option)
		if err != nil {
			log.Printf("TenantMiddleware:parseTenantAccount error:%s.\n", err.Error())
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		var identity = &TenantIdentity{}

		cache, ok := GetRedisCacheFromContext(c)
		if !ok {
			// redis instance is down.
			log.Printf("TenantMiddleware:GetRedisCacheFromContext error:%s.\n", ErrNotFoundCache.Error())
			identity, err = loadTenantAccount(c, account)
			if err != nil {
				log.Printf("TenantMiddleware:loadTenantAccount error:%s\n", err.Error())
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.Set(tenantContextKey, identity)
			c.Next()
			return
		}

		var json string
		json, err = cache.Get(account).Result()
		if err != nil {
			log.Printf("TenantMiddleware:cache.Get error:%s\n", err.Error())

			identity, err = loadTenantAccount(c, account)
			if err != nil {
				log.Printf("TenantMiddleware:loadTenantAccount error:%s\n", err.Error())
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.Set(tenantContextKey, identity)

			if err == redis.Nil {
				// redis has no record
				err = cacheTenantAccount(cache, identity, account)
				if err != nil {
					log.Printf("TenantMiddleware:cacheTenantAccount error:%s\n", err.Error())
				}
			}

			c.Next()
			return
		}

		if err := jsoniter.Unmarshal([]byte(json), identity); err != nil {
			log.Printf("TenantMiddleware:jsoniter.Unmarshal error:%s\n", err.Error())
		}
		c.Set(tenantContextKey, identity)
		c.Next()
	}
}

//GetTenantIdentityFromContext return one TenantIdentity instance or nil from current context.
func GetTenantIdentityFromContext(c *gin.Context) (*TenantIdentity, bool) {
	value, exist := c.Get(tenantContextKey)
	if !exist {
		return nil, false
	}

	identity, ok := value.(*TenantIdentity)
	if !ok {
		return nil, false
	}

	return identity, true
}

func parseTenantAccount(c *gin.Context, option *TenantOption) (string, error) {
	if option == nil {
		return "", errMissingOption
	}

	var account string
	if strings.ToLower(option.Mode) == "domain" {
		account = parseSubdomain(c.Request.Host, option.Host)
	} else if strings.ToLower(option.Mode) == "header" {
		account = parseHeader(c.Request.Header, option.Key)
	} else if strings.ToLower(option.Mode) == "url" {
		account = parseURL(c, option.Key)
	} else {
		return "", errNotImplementTenantMode
	}

	if account == "" || utils.ContainsString(option.BlackList, account) {
		return "", errInvalidTenantAccount
	}
	return account, nil
}

func parseSubdomain(requestHost, appHost string) string {
	return strings.Replace(
		strings.ToLower(requestHost),
		"."+strings.ToLower(appHost), "", -1)
}

func parseHeader(header http.Header, tenantKey string) string {
	if strings.Trim(tenantKey, "") == "" {
		return ""
	}
	return strings.ToLower(header.Get(tenantKey))
}

func parseURL(c *gin.Context, tenantKey string) string {
	if strings.Trim(tenantKey, "") == "" {
		return ""
	}
	return strings.ToLower(c.Param(tenantKey))
}

func loadTenantAccount(c *gin.Context, account string) (*TenantIdentity, error) {
	var identity = &TenantIdentity{}
	db, ok := GetReadonlyDBFromContext(c)
	if !ok {
		return nil, ErrNotFoundReadonlyDB
	}

	row := db.QueryRow(`SELECT id, account FROM account.tenants WHERE LOWER(account)=$1 AND enabled=true;`, account)
	err := row.Scan(&identity.ID, &identity.Account)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errInvalidTenantAccount
		}
		return nil, err
	}
	return identity, nil
}

func cacheTenantAccount(cache *redis.Client, identity *TenantIdentity, account string) error {
	json, err := jsoniter.Marshal(identity)
	if err != nil {
		return err
	}

	err = cache.Set(account, json, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
