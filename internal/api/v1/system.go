package v1

import (
	"github.com/gin-gonic/gin"
	base "github.com/sqc157400661/helper/api"
	"github.com/sqc157400661/kdb-sidecar/pkg/mysql/config"
)

type System struct {
	base.Api
}

func (e System) PodInfo(c *gin.Context) {
	e.OK(map[string]string{
		"podName": config.Hostname,
		"podIp":   config.PodIP,
	}, "查询成功")
}
