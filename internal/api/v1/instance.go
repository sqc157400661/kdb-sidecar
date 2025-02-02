package v1

import (
	"github.com/gin-gonic/gin"
	base "github.com/sqc157400661/helper/api"
	"github.com/sqc157400661/kdb-sidecar/internal/biz"
	"github.com/sqc157400661/kdb-sidecar/internal/types"
)

type Instance struct {
	base.Api
}

// GetPage
// @Summary 获取实例列表
// @Description 获取实例列表
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
	list, err := biz.ListInstance(req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(list, "查询成功")
}
