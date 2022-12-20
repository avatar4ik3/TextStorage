package textHandlers

import (
	"avatar4ik3/TextStorage/api/handlers"
	"avatar4ik3/TextStorage/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RemoveTextByIdHandler struct {
	logger *logrus.Logger
	repo   *models.Repository
}

func NewRemoveTextByIdHandler(logger *logrus.Logger, repo *models.Repository) *RemoveTextByIdHandler {
	return &RemoveTextByIdHandler{
		logger: logger,
		repo:   repo,
	}
}

func (this *RemoveTextByIdHandler) Handle() *handlers.Handler {
	return &handlers.Handler{
		Path:   "/texts/:id",
		Method: http.MethodDelete,
		Func: func(ctx *gin.Context) {

			id := handlers.TryWithErrorG(
				func() (uint64, error) {
					id := uint64(0)
					return id, ctx.BindUri(&id)
				},
				http.StatusBadRequest,
				ctx,
			)
			handlers.TryWithError(
				func() error {
					return this.repo.RemoveText(id)
				},
				http.StatusBadRequest,
				ctx,
			)
			if len(ctx.Errors) == 0 {
				ctx.JSON(
					http.StatusOK,
					gin.H{
						"succeed": true,
					},
				)
			}
		},
	}
}
