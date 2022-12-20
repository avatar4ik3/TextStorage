package textHandlers

import (
	"avatar4ik3/TextStorage/api/handlers"
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
	Description string `json:"description"`
	Value       string `json:"value"`
}

func NewAddTextHandler(logger *logrus.Logger, repo *models.Repository) *AddTextHandler {
	return &AddTextHandler{
		logger: logger,
		repo:   repo,
	}
}

func (this *AddTextHandler) Handle() *handlers.Handler {
	return &handlers.Handler{
		Path:   "/texts",
		Method: http.MethodPost,
		Func: func(ctx *gin.Context) {

			req := handlers.TryWithErrorG(
				func() (AddTextRequest, error) {
					req := AddTextRequest{}
					return req, ctx.BindJSON(&req)
				},
				http.StatusBadRequest,
				ctx,
			)
			this.logger.Info(req)
			text := handlers.TryWithErrorG(func() (*models.Text, error) {
				return this.repo.AddText(req.Value, req.Description)
			}, http.StatusInternalServerError, ctx)
			if len(ctx.Errors) == 0 {
				ctx.JSON(
					http.StatusCreated,
					text,
				)
			}
		},
	}
}
