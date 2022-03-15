package util

import "github.com/gin-gonic/gin"

type Gin struct {
	Ctx *gin.Context
}

type Json struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (g *Gin) Json(status int, message string, data interface{}) {
	g.Ctx.JSON(200, Json{Status: status, Message: message, Data: data})
	return
}
