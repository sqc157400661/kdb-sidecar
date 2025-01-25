package middleware

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sqc157400661/helper/api/common"
	"io"
	"net/http"
	"time"
)

// ReqLogger 记录请求日志
func ReqLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		var body string
		switch c.Request.Method {
		case http.MethodPost, http.MethodPut, http.MethodGet, http.MethodDelete:
			bf := bytes.NewBuffer(nil)
			wt := bufio.NewWriter(bf)
			_, err := io.Copy(wt, c.Request.Body)
			if err != nil {
				fmt.Printf("copy body error, %s \n", err.Error())
				err = nil
			}
			rb, _ := io.ReadAll(bf)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(rb))
			body = string(rb)
		}

		// 处理请求
		c.Next()

		// 预检的跳过
		if c.Request.Method == http.MethodOptions {
			return
		}

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime).Seconds()

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		rt, bl := c.Get("result")
		var result = ""
		if bl {
			rb, err := json.Marshal(rt)
			if err != nil {
				fmt.Printf("json Marshal req result error, %s \n", err.Error())
			} else {
				result = string(rb)
			}
		}

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		common.Logger(c).Infow("API Record",
			"status_code", statusCode,
			"latency_time", latencyTime,
			"client_ip", clientIP,
			"req_method", reqMethod,
			"body", body,
			"req_uri", reqUri,
			"result", result,
		)
	}
}
