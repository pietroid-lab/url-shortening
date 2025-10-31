package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateURLRequest struct {
	OriginalURL string `json:"originalUrl" binding:"required,url"`
	CustomSlug  string `json:"customSlug,omitempty"`
}

type Service interface {
	GenerateShorterURL(ctx context.Context, originalURL, slug string) (string, error)
}

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) GenerateURL(c *gin.Context) {
	ctx := c.Request.Context()

	var req CreateURLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	url, err := h.service.GenerateShorterURL(ctx, req.OriginalURL, req.CustomSlug)	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate shorter URL",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "URL received",
		"url":     url,
		"slug":    req.CustomSlug,
	})
}
