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

// GetAlbums is a method on Handler and has the signature compatible with
// gin.HandlerFunc (func(*gin.Context)) because the receiver is bound when
// passing the method value (h.GetAlbums) to Gin.
func (h *Handler) GetAlbums(c *gin.Context) {
	// Use h.service when you need to call business logic. For now return the
	// api.Url as previously implemented.
	c.IndentedJSON(http.StatusOK, api.Url)
}
