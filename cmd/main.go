package main

import (
	"url-shorting/cmd/api/handler"
	"url-shorting/cmd/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	repo := &internal.Repository{}
	svc := internal.NewService(repo)
	h := handler.NewHandler(svc)

	router.GET("/URLs", h.GetAlbums)

	router.Run("localhost:8080")
}
