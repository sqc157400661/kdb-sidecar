package server

import (
	"fmt"
	"github.com/sqc157400661/helper/conf"
	"github.com/sqc157400661/kdb-sidecar/internal/api/router"
	"net/http"
	"time"
)

func StartServer() (err error) {
	//路由配置
	routersInit := router.InitRouter()
	//http服务器配置
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", conf.GetIntD("server.port", 8080)),
		Handler:        routersInit,
		ReadTimeout:    conf.GetDurationD("server.readTimeout", 30) * time.Millisecond,
		WriteTimeout:   conf.GetDurationD("server.writeTimeout", 30) * time.Millisecond,
		MaxHeaderBytes: 1 << 20,
	}

	//启动服务器
	err = s.ListenAndServe()
	return
}
