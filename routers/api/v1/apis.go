package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jihuago/SSO/controllers"
)

func LoadApiRouter(e *gin.Engine) {
	e.GET("/", controllers.Index)
}
