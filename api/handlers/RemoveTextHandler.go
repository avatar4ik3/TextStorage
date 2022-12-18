package handlers

import (
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
func (this *RemoveTextHandler) Handle() *Handler {
	return &Handler{
		Path:   "/texts",
		Method: http.MethodDelete,
		Func: func(ctx *gin.Context) {
			req := &RemoveRequest{}
			this.logger.Info("Recieved remove texts!")
			if err := ctx.BindJSON(&req); err != nil {
				ctx.JSON(
					http.StatusBadRequest,
					gin.H{
						"error": err.Error(),
					},
				)
			}
			responses := []RemoveResponse{}
			for _, id := range req.Ids {
				r := &RemoveResponse{Id: id}
				if err := this.repo.RemoveText(id); err != nil {
					r.Error = err.Error()
				}
				responses = append(responses, *r)
			}
			ctx.JSON(
				http.StatusOK,
				responses,
			)
		},
	}

}
