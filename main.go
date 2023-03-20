package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//hello
	r := gin.Default()

	//Grouping for covid API
	//covid_api := r.Group("/covid")
	{
		//Endpoint for Getting Covid Summary
		//covid_api.GET("/summary", controller.SummaryController)
	}

	//Check for wrong route
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"err_response": "404 Not Found (Are you sure you typed /covid/summary and GET method ?)",
		})
	})

	r.Run(":8080")
}
