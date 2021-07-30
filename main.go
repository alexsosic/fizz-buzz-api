package main

import (
	"github.com/alexsosic/fizz-buzz-api/controllers"
	"github.com/alexsosic/fizz-buzz-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/api/:int1/:int2/:limit/:str1/:str2", controllers.FizzBuzz)
	r.GET("/", controllers.GetStats)

	// Run the server
	r.Run()
}
