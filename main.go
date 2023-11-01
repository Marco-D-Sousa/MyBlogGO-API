package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"myblog/controllers"
	"myblog/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostCreate)

	log.Fatal(r.Run())
}
