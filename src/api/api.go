package api

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/khodealib/golang-clean-web-api/src/api/routers"
	"github.com/khodealib/golang-clean-web-api/src/api/validations"
	"github.com/khodealib/golang-clean-web-api/src/config"
)

func InitServer() {
	cfg := config.GetConfig()

	if cfg.Server.RunMode != "" && (cfg.Server.RunMode == "debug" || cfg.Server.RunMode == "release") {
		gin.SetMode(cfg.Server.RunMode)
	} else {
		panic(fmt.Sprintf("Invalid run mode: %s", cfg.Server.RunMode))
	}

	validations.RegisterCustomValidations()

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		routers.HealthRouter(health)
	}

	err := r.Run(fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port))
	if err != nil {
		panic(err)
	}
}
