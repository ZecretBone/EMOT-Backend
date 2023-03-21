package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"penguin.com/emot/Controller"
)

func main() {
	//hello
	r := gin.Default()

	//Grouping for covid API
	emot_api := r.Group("/back-emot")
	{
		//Endpoint for Getting Covid Summary
		emot_api.GET("/newrec", Controller.UserController)
		emot_api.GET("/newuser", Controller.UserController)
	}

	//Check for wrong route
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"err_response": "404 Not Found (Are you sure you typed /covid/summary and GET method ?)",
		})
	})

	r.Run(":8080")
}
