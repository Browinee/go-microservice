package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func getPageInfo(c *gin.Context) (offset, limit int) {
	offsetStr := c.DefaultQuery("pn", "0")
	limitStr := c.DefaultQuery("psize", "10")
	var (
		err error
	)
 	offset, err = strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}
	return offset, limit
}
