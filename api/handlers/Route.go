package handlers

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Path   string
	Method string
	Func   gin.HandlerFunc
}
type Route interface {
	Handle() *Handler
}
