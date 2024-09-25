package handler

import (
	"github.com/gin-gonic/gin"
)

//retrieve path variable
func getIDParam(c *gin.Context) string {
	return c.Param("id")
}
