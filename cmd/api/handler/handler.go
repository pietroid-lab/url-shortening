package handler

import (
	"net/http"
	"url-shorting/cmd/api"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, api.Url)

}
