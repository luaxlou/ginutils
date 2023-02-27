package ginutils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetInt64Param(c gin.Context, key string) int64 {

	str := c.Param(key)

	num, _ := strconv.ParseInt(str, 10, 64)

	return num

}

func GetIdParam(c gin.Context, key string) int64 {

	return GetInt64Param(c, "id")

}
