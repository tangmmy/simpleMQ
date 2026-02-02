package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"simpleMQ/internal/handlers"
)

func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		handlers.HandleTest(c)
	})

	r.POST("/produce", func(c *gin.Context) {
		handlers.HandleProducerMessage(c)
	})

	r.POST("/consume", func(c *gin.Context) {
		handlers.HandleMessage(c)
	})

	r.POST("/publish", func(c *gin.Context) {
		handlers.HandleMessage(c)
	})

	r.POST("/subscribe", func(c *gin.Context) {
		handlers.HandleMessage(c)
	})

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
