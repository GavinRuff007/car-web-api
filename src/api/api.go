package api

import (
	"fmt"

	"github.com/GavinRuff007/car-web-api/api/routers"
	"github.com/GavinRuff007/car-web-api/config"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		health := v1.Group("/health")
		routers.Health(health)
	}
	r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
}
