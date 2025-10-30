package handler

import (
	"net/http"
	"url-shorting/cmd/api"
	"url-shorting/cmd/internal"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *internal.Service
}

func NewHandler(s *internal.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, api.Url)
}
