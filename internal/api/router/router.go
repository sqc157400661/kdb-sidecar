package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sqc157400661/kdb-sidecar/internal/api/middleware"
	v1 "github.com/sqc157400661/kdb-sidecar/internal/api/v1"
)

func InitRouter() *gin.Engine {

	//设置gin运行模式
	//gin.SetMode(conf.GetString("runMode"))

	r := gin.New()
	r.Use(middleware.Cors())
	gin.SetMode(gin.DebugMode)
	r.Any("health", func(c *gin.Context) {
		c.JSON(200, "ok")
	})
	//恢复中间件 默认可用：r.Use(gin.Recovery())
	r.Use(gin.Recovery())
	r.Use(middleware.Trace())
	//设置路由分组
	NoCheckRoleRouter(r)
	//CheckRoleRouter(r, middleware.Auth(buSvcCli), middleware.ReqLogger())
	return r
}

// NoCheckRoleRouter 不需要认证的路由
func NoCheckRoleRouter(r *gin.Engine) {
	v1Group := r.Group("/v1")
	monitorRouter := v1Group.Group("/mysql")
	{
		m := v1.Instance{}
		monitorRouter.GET("/instance", m.GetPage)
	}
}
