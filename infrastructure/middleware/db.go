package middleware

import (
	"errors"

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
	instance, err := db.GetDefaultDB()
	if err != nil {
		panic(err)
	}
	return func(c *gin.Context) {
		c.Set(vandadb, instance)
		c.Next()
	}
}

//ReadonlyDBMiddleware adds readonly db instance into context.
func ReadonlyDBMiddleware() gin.HandlerFunc {
	instance, err := db.GetReadonlyDB()
	if err != nil {
		panic(err)
	}
	return func(c *gin.Context) {
		c.Set(vandadb, instance)
		c.Next()
	}
}

// GetDefaultDB return one instance from context.
func GetDefaultDB(c *gin.Context) (db.DefaultDB, bool) {
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

// GetReadonlyDB return one instance from context.
func GetReadonlyDB(c *gin.Context) (db.ReadonlyDB, bool) {
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
