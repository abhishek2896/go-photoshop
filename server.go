// Package main is the CLI.
// You can use the CLI via Terminal.
package main

import (
	// "fmt"
	"github.com/Abhishek-Nagarkoti/go-photoshop/handler"
	"github.com/Abhishek-Nagarkoti/go-photoshop/middleware"
	"github.com/gin-gonic/gin"
	"os"
)

const (
	// Port at which the server starts listening
	Port = "8000"
)

func main() {

	// Configure
	router := gin.Default()
	//using cors
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, X-HTTP-Method-Override,Authorization, Content-Type, Accept")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		if c.Request.Method == "OPTIONS" {
			c.JSON(200, gin.H{"All": "Good"})
		} else {
			c.Next()
		}

	})
	// Middlewares
	router.Use(middlewares.ErrorHandler)

	v1 := router.Group("/api")
	{
		v1.POST("/image", handler.Create)
		v1.PUT("/image/:image", handler.Update)
		v1.GET("/image/:image", handler.Get)
	}

	// Start listening
	port := Port
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}
	router.Run(":" + port)
}
