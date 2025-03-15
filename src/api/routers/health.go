package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/khodealib/golang-clean-web-api/src/api/handlers"
)

func HealthRouter(r *gin.RouterGroup) {
	healthHandler := handlers.NewHealthHandler()
	r.GET("/", healthHandler.Health)
}
