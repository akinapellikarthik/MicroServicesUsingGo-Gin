package main

import (
	"GIN-DEMO/database"
	"GIN-DEMO/middlewares"
	"github.com/gin-gonic/gin"

	"GIN-DEMO/controller"
)

func main() {

	database.Connect("root:admin123@tcp(localhost:3306)/test_db?parseTime=true")
	database.Migrate()
	r := initRouter()
	r.GET("/hello", controller.HelloWorld)
	r.Run(":8989")

}

func initRouter() *gin.Engine {
	router := gin.Default()
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		return nil
	}
	api := router.Group("/api")
	{
		api.POST("/token", controller.GenerateToken)
		api.POST("/user/register", controller.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controller.Ping)
		}
	}
	return router
}
