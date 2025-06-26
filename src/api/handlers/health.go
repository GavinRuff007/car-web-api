package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Header struct {
	UserId    string
	BrowserId string
}

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Fullname": "Parsa Eftekharmanesh",
	})
}

func (h *HealthHandler) HealthParam(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, fmt.Sprintf("Working! %s", id))
}

func (h *HealthHandler) HealthTest(c *gin.Context) {
	header := Header{}
	c.BindHeader(&header)
	c.JSON(http.StatusOK, gin.H{
		"id":       header,
		"FullName": "Parsa Eftekharmanesh",
	})
}

func (h *HealthHandler) HealthQuery(c *gin.Context) {
	id := c.Query("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *HealthHandler) HealthBody(c *gin.Context) {
	p := ModelTest{}
	c.ShouldBindJSON(&p)
	c.JSON(http.StatusOK, gin.H{
		"id": p,
	})
}
