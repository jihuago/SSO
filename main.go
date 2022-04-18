package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jihuago/SSO/routers"
	"github.com/spf13/viper"
)

func main() {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	bootstrap()

	routers.SetupRouter(r)

	port := viper.GetInt("app.port")

	r.Run(fmt.Sprintf(":%d", port))
}

func bootstrap() {
	// 加载配置组件
	viper.SetConfigFile("./conf/app.yaml")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}
}
