package main

import (
	"url-shorting/cmd/api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/URLs", handler.GetAlbums)

	router.Run("localhost:8080")
}
