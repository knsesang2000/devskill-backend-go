package handlers

import "github.com/gin-gonic/gin"

// Router registers handler routes under /api/v1
func Router(r *gin.Engine) {
	api := r.Group("/api/v1")
	api.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })
	RegisterRoutes(api)
}
