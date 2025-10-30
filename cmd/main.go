package main

import (
	"url-shorting/cmd/api/handler"
	"url-shorting/cmd/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Create repository and service, then handler. The handler's method
	// value (h.GetAlbums) has type func(*gin.Context), so it can be used
	// directly as a gin handler.
	repo := &internal.Repository{}
	svc := internal.NewService(repo)
	h := handler.NewHandler(svc)

	router.GET("/URLs", h.GetAlbums)

	router.Run("localhost:8080")
}
