package v1

import (
	"github.com/gin-gonic/gin"
	base "github.com/sqc157400661/helper/api"
	"github.com/sqc157400661/kdb-sidecar/internal/types"
)

type Instance struct {
	base.Api
}

// GetPage
// @Summary 获取集群节点拓扑列表
// @Description 获取集群节点拓扑列表
// @Tags GetInstanceList
// @ID get-instance-list
// @Param status query string false "status"
// @Success 200 {object} response.DefaultResponse "{"code": 200, "data": [...]}"
// @Router /api/v1/instance [get]
func (e Instance) GetPage(c *gin.Context) {
	req := types.InstancesReq{}
	err := e.MakeContext(c).
		Bind(&req).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(nil, "查询成功")
}
