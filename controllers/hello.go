package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func Index(ctx *gin.Context) {
	appName := viper.Get("app.port")
	ctx.JSON(http.StatusOK, gin.H{
		"message": appName,
	})
}
