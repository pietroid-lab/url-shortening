package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/URLs", getAlbums)

	router.Run("localhost:8080")
}