package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateExpense(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Status" : "SUCCESS",
		"Message": "Updated Successfully",
	})
}
