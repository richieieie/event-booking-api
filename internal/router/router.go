package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	routes "github.com/richieieie/event-booking/internal/router/v1"
)

func NewGinRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	// Set up middlewares

	// Set up routes
	apiV1 := r.Group("/api/v1")
	routes.InitEventHandler(apiV1)
	routes.InitUserHandler(apiV1)

	return r
}
