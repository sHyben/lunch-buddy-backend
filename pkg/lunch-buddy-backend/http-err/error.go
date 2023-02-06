package http_err

import "github.com/gin-gonic/gin"

// NewError example function
// @Summary Error
// @Description Error
// @Produce json
// @Success 400 {object} HTTPError
// @Router / [get]
// @Security Authorization Token
func NewError(c *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	c.JSON(status, er)
}

// HTTPError example
// @Summary Error
// @Description Error
// @Produce json
// @Success 400 {object} HTTPError
// @Router / [get]
// @Security Authorization Token
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
