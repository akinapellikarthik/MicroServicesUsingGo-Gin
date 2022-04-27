package controller

import (
	"GIN-DEMO/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//TODO: Need to write test cases for service class
func HelloWorld(ctx *gin.Context) {


	msg := service.DemoService.HelloService(ctx.Param("cust"))

	ctx.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}
