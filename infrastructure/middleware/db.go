package middleware

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/allenslian/vanda/infrastructure/factory/db"
)

var (
	//ErrNotFoundDefaultDB means context has not default db instance.
	ErrNotFoundDefaultDB = errors.New("Don't find default db instance in the current context")
	//ErrNotFoundReadonlyDB means context has not default db instance.
	ErrNotFoundReadonlyDB = errors.New("Don't find readonly db instance in the current context")
)

const (
	vandadb = "vanda_db"
)

//DefaultDBMiddleware adds default db instance into context.
func DefaultDBMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		instance, err := db.GetDefaultDB()
		if err != nil {
			log.Printf("DefaultDBMiddleware:GetDefaultDB error:%s.\n", err.Error())
			c.Next()
			return
		}
		c.Set(vandadb, instance)
		c.Next()
	}
}

//ReadonlyDBMiddleware adds readonly db instance into context.
func ReadonlyDBMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		instance, err := db.GetReadonlyDB()
		if err != nil {
			log.Printf("ReadonlyDBMiddleware:GetReadonlyDB error:%s.\n", err.Error())
			c.Next()
			return
		}

		c.Set(vandadb, instance)
		c.Next()
	}
}

// GetDefaultDBFromContext return one instance from context.
func GetDefaultDBFromContext(c *gin.Context) (db.DefaultDB, bool) {
	value, exist := c.Get(vandadb)
	if !exist {
		return nil, false
	}
	vandaDB, ok := value.(db.DefaultDB)
	if !ok {
		return nil, false
	}
	return vandaDB, true
}

// GetReadonlyDBFromContext return one instance from context.
func GetReadonlyDBFromContext(c *gin.Context) (db.ReadonlyDB, bool) {
	value, exist := c.Get(vandadb)
	if !exist {
		return nil, false
	}
	vandaDB, ok := value.(db.ReadonlyDB)
	if !ok {
		return nil, false
	}
	return vandaDB, true
}
