package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ErrorHandler(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, ginErr := range c.Errors {
			logger.WithError(ginErr.Err).Error(ginErr.Err.Error())
		}
		c.JSON(int(c.Errors.Last().Meta.(int)), gin.H{
			"errors": c.Errors.Last().Err.Error(),
		} /* error payload */)
	}
}
