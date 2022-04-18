package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/jihuago/SSO/routers/api/v1"
)

// 加载路由
func SetupRouter(engine *gin.Engine) {
	v1.LoadApiRouter(engine)
}
