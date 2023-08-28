package config

import (
	"github.com/gin-gonic/gin"
)

func InitEndpoints() {
	router := gin.Default()

	router.Use(func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Next()
	})

	router.Run("localhost:8080")
}
