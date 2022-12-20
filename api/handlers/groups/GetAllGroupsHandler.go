package groupsHandlers

import (
	"avatar4ik3/TextStorage/api/handlers"
	"avatar4ik3/TextStorage/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type GetAllGroupsHandler struct {
	logger *logrus.Logger
	repo   *models.Repository
}

func NewGetAllGroupsHandler(logger *logrus.Logger,
	repo *models.Repository) *GetAllGroupsHandler {
	return &GetAllGroupsHandler{
		logger: logger,
		repo:   repo,
	}
}

func (this *GetAllGroupsHandler) Handle() *handlers.Handler {
	return &handlers.Handler{
		Path:   "/groups",
		Method: http.MethodGet,
		Func: func(ctx *gin.Context) {
			data := handlers.TryWithErrorG(func() ([]models.Group, error) {
				return this.repo.AllGroups()
			}, http.StatusInternalServerError, ctx)
			if len(ctx.Errors) == 0 {
				ctx.JSON(
					http.StatusOK,
					gin.H{
						"groups": data,
					},
				)
			}
		},
	}
}
