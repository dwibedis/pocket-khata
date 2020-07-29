package resource

import (
	"github.com/gin-gonic/gin"
	"log"
)

func InitServer() {
	r := gin.Default()
	//r.GET("/", controller.CreateExpense)
	err := r.Run("localhost:8080")
	if err != nil {
		log.Panic("Server up failed!!")
	}
}
