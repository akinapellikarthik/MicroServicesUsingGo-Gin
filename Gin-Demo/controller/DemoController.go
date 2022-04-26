package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//TODO: Need to write test cases for service class
func HelloWorld(ctx *gin.Context) {


	//msg := service.DemoServiceImpl{}.HelloService(ctx.Query("name"))

	ctx.JSON(http.StatusOK, gin.H{
		"message": "test",
	})
}
