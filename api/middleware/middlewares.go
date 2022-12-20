package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewErrorHandler(logger *logrus.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		for _, ginErr := range ctx.Errors {
			logger.WithError(ginErr.Err).Error(ginErr.Err.Error())
		}
		if ctx.Errors.Last() != nil {
			ctx.JSON(-1, gin.H{
				"errors": ctx.Errors.Last().Err.Error(),
			})
		}
	}
}
