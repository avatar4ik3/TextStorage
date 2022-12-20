package textHandlers

import (
	"avatar4ik3/TextStorage/api/handlers"
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
func (this *GetAllTextsHandler) Handle() *handlers.Handler {
	return &handlers.Handler{
		Path:   "/texts",
		Method: http.MethodGet,
		Func: func(ctx *gin.Context) {
			this.logger.Info("Recieved get all text!")
			res := handlers.TryWithErrorG(func() ([]models.Text, error) {
				return this.repo.AllTexts()
			}, http.StatusInternalServerError, ctx)
			ctx.JSON(
				http.StatusOK,
				res,
			)
		},
	}

}
