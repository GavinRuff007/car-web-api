package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(200, "Working!")
}

func (h *HealthHandler) HealthId(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, fmt.Sprintf("Working! %s", id))

}
