package main

import (
	"example/webservices/controllers"
	"example/webservices/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/create", controllers.TaskCreate)
	r.GET("/getOne/:id", controllers.GetOne)
	r.GET("/getAll", controllers.GetAll)
	r.PUT("/update/:id", controllers.Update)
	r.DELETE("/delete/:id", controllers.Delete)

	r.Run() // listen and serve on 0.0.0.0:8080
}

