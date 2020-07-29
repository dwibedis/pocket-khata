package main

import (
	"github.com/dwibedis/pocket-khata/controllers"
	"github.com/gin-gonic/gin"
	"log"
)

func InitServer() {
	r := gin.Default()
	r.GET("/", controllers.CreateExpenses)
	err := r.Run("localhost:8080")
	if err != nil {
		log.Panic("Server up failed!!")
	}
}
