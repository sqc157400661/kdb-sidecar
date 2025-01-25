package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sqc157400661/helper/api/response"
)

// Trace 日志记录到文件
func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqesutId string
		var has bool
		reqesutId, has = c.GetQuery("RequestId")
		if !has {
			reqesutId, has = c.GetQuery("requestId")
		}
		if has {
			c.Request.Header.Set(response.TrafficKey, reqesutId)
		}
	}
}
