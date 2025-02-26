package api

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/khodealib/golang-clean-web-api/src/api/routers"
	"github.com/khodealib/golang-clean-web-api/src/config"
)

func InitServer() {
	config := config.GetConfig()
	fmt.Printf("Config %+v\n", config)
	if config.Server.RunMode == "debug" {
		gin.SetMode(gin.DebugMode)
	} else if config.Server.RunMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		panic(fmt.Sprintf("Invalid run mode: %s", config.Server.RunMode))
	}

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		health := v1.Group("/health")
		routers.HealthRouter(health)
	}

	err := r.Run(fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port))
	if err != nil {
		panic(err)
	}
}
