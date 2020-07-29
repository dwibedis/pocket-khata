package main

import (
	"github.com/gin-gonic/gin"

	"log"
)

func main() {
	r := gin.Default()
	err := r.Run("localhost:8080")
	if err != nil {
		log.Panic("Server up failed!!")
	}
}
