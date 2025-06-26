package routers

import (
	"github.com/GavinRuff007/car-web-api/api/handlers"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()
	r.GET("/", handler.Health)
	r.POST("/:id", handler.HealthParam)
	r.POST("/test", handler.HealthTest)
	r.POST("/ds", handler.HealthQuery)
	r.POST("/dsw", handler.HealthBody)

}
