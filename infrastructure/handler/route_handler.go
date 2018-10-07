package handler

import (
	"github.com/gin-gonic/gin"
)

const (
	// Web means it's one web handler.
	Web Kind = iota

	// API means it's one api handler.
	API
)

type (
	// Kind indicates one handler is api handler or web handler.
	Kind int

	// IRouteHandler describes one route handler.
	IRouteHandler interface {
		GetKind() Kind
		GetURL() string
		GetMethod() string
		GetHandlers() []gin.HandlerFunc
	}
)

// NewRoute initializes one IRouteHandler instance.
func NewRoute(kind Kind, url, method string,
	handler gin.HandlerFunc,
	middlewares ...gin.HandlerFunc) IRouteHandler {
	h := &routeHandler{
		kind:    kind,
		url:     url,
		method:  method,
		handler: handler,
	}

	for _, mw := range middlewares {
		h.middlewares = append(h.middlewares, mw)
	}
	return h
}

// inner implement of IRouteHandler interface.
type routeHandler struct {
	kind        Kind
	url         string
	method      string
	middlewares []gin.HandlerFunc
	handler     gin.HandlerFunc
}

func (h *routeHandler) GetKind() Kind {
	return h.kind
}

func (h *routeHandler) GetURL() string {
	return h.url
}

func (h *routeHandler) GetMethod() string {
	return h.method
}

func (h *routeHandler) GetHandlers() []gin.HandlerFunc {
	return append(h.middlewares, h.handler)
}
