package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type PingPongHandler struct {
	logger *logrus.Logger
}

func NewPingPongHandler(logger *logrus.Logger) *PingPongHandler {
	return &PingPongHandler{
		logger: logger,
	}
}

func (this *PingPongHandler) Handle() *Handler {
	return &Handler{
		Path:   "/ping",
		Method: http.MethodGet,
		Func: 	func(ctx *gin.Context) {
					this.logger.Info("Recieved ping!")
					ctx.JSON(
						http.StatusOK,
						gin.H{
							"message": "pong",
						},
					)
				},
	}
}
