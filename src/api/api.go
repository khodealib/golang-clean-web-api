package api

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/khodealib/golang-clean-web-api/src/api/routers"
	"github.com/khodealib/golang-clean-web-api/src/config"
)

func InitServer() {
	config := config.GetConfig()

	if config.Server.RunMode != "" && (config.Server.RunMode == "debug" || config.Server.RunMode == "release") {
		gin.SetMode(config.Server.RunMode)
	} else {
		panic(fmt.Sprintf("Invalid run mode: %s", config.Server.RunMode))
	}

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		routers.HealthRouter(health)
	}

	err := r.Run(fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port))
	if err != nil {
		panic(err)
	}
}
