package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetParamID 获取param里的Int64 ID
func GetParamID(c *gin.Context) int64 {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	CheckError(err)
	return id
}
