package middlewares

import "github.com/gin-gonic/gin"

// NoMethodHandler godoc
// @Summary Method not allowed
// @Description Method not allowed
// @Produce json
// @Success 405 {object} gin.H
// @Router / [get]
// @Security Authorization Token
func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(405, gin.H{"message": "Method not allowed"})
	}
}

// NoRouteHandler godoc
// @Summary Route not found
// @Description Route not found
// @Produce json
// @Success 404 {object} gin.H
// @Router / [get]
// @Security Authorization Token
func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "The processing function of the request route was not found"})
	}
}
