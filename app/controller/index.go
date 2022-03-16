package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "index/index.html", gin.H{
		"msg": "hello gofun",
	})
}
