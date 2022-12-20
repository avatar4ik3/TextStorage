package textHandlers

import (
	"avatar4ik3/TextStorage/api/handlers"
	models "avatar4ik3/TextStorage/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RemoveTextHandler struct {
	logger *logrus.Logger
	repo   *models.Repository
}

type RemoveRequest struct {
	Ids []uint64 `json:"ids"`
}

type RemoveResponse struct {
	Id    uint64 `json:"id"`
	Error string `json:"error"`
}

func NewRemoveTextHandler(logger *logrus.Logger, repo *models.Repository) *RemoveTextHandler {
	return &RemoveTextHandler{
		logger: logger,
		repo:   repo,
	}
}
func (this *RemoveTextHandler) Handle() *handlers.Handler {
	return &handlers.Handler{
		Path:   "/texts",
		Method: http.MethodDelete,
		Func: func(ctx *gin.Context) {
			req := handlers.TryWithErrorG(func() (RemoveRequest, error) {
				req := RemoveRequest{}
				return req, ctx.BindJSON(&req)
			}, http.StatusBadRequest, ctx)

			responses := []RemoveResponse{}
			for _, id := range req.Ids {
				handlers.TryWithError(func() error {
					return this.repo.RemoveText(id)
				}, http.StatusInternalServerError, ctx)
			}
			ctx.JSON(
				http.StatusOK,
				responses,
			)
		},
	}

}
