package handlers

import (
	models "avatar4ik3/TextStorage/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AddTextHandler struct {
	logger *logrus.Logger
	repo   *models.Repository
}

type AddTextRequest struct {
	Value string `json:"value"`
}

func NewAddTextHandler(logger *logrus.Logger, repo *models.Repository) *AddTextHandler {
	return &AddTextHandler{
		logger: logger,
		repo:   repo,
	}
}

func (this *AddTextHandler) Handle() *Handler {
	return &Handler{
		Path:   "/texts",
		Method: http.MethodPost,
		Func: func(ctx *gin.Context) {
			this.logger.Info("Recieved add text!")

			req := &AddTextRequest{}
			if err := ctx.BindJSON(&req); err != nil {
				ctx.JSON(
					http.StatusBadRequest,
					gin.H{
						"error": err,
					},
				)
			}
			this.logger.Info(req.Value)
			text := &models.Text{
				Value: req.Value,
			}
			this.repo.AddText(text)
			ctx.JSON(
				http.StatusOK,
				text,
			)
		},
	}
}
