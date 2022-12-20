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

func TryWithErrorG[T any](f func() (T, error), status int, ctx *gin.Context) T {
	res, err := f()
	if err != nil {
		ctx.Errors = append(ctx.Errors, &gin.Error{Err: err, Meta: status})
	}
	return res
}

func TryWithError(f func() error, status int, ctx *gin.Context) {
	err := f()
	if err != nil {
		ctx.Errors = append(ctx.Errors, &gin.Error{Err: err, Meta: status})
	}
}
