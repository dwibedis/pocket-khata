package main

import (
	"github.com/gin-gonic/gin"
	//"github.com/dwibedis/controllers"
	"log"
)

func main() {
	r := gin.Default()
	//r.GET("/", controllers.)
	err := r.Run("localhost:8080")
	if err != nil {
		log.Panic("Server up failed!!")
	}
}
