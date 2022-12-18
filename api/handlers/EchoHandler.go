package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type EchoHandler struct {
	logger *logrus.Logger
}

func NewEchoHandler(logger *logrus.Logger) *EchoHandler {
	return &EchoHandler{
		logger: logger,
	}
}

func (this *EchoHandler) Handle() *Handler {
	return &Handler{
		Path:   "/echo",
		Method: http.MethodGet,
		Func: func(ctx *gin.Context) {
			this.logger.Info("Recieved echo!")
			ctx.JSON(
				http.StatusOK,
				gin.H{
					"message": time.Now(),
				},
			)
		},
	}
}
