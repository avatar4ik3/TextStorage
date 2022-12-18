package handlers

import (
	models "avatar4ik3/TextStorage/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type GetAllTextsHandler struct {
	logger *logrus.Logger
	repo   *models.Repository
}

func NewGetAllTextsHandler(logger *logrus.Logger, repo *models.Repository) *GetAllTextsHandler {
	return &GetAllTextsHandler{
		logger: logger,
		repo:   repo,
	}
}
func (this *GetAllTextsHandler) Handle() *Handler {
	return &Handler{
		Path:   "/texts",
		Method: http.MethodGet,
		Func: func(ctx *gin.Context) {
			this.logger.Info("Recieved get all text!")
			res, err := this.repo.AllTexts()
			if err != nil {
				ctx.JSON(
					http.StatusBadRequest,
					gin.H{
						"error": err,
					},
				)
			}
			ctx.JSON(
				http.StatusOK,
				res,
			)
		},
	}

}
