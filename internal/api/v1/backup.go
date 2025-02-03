package v1

import (
	"github.com/gin-gonic/gin"
	base "github.com/sqc157400661/helper/api"
	"github.com/sqc157400661/kdb-sidecar/internal/biz"
	"github.com/sqc157400661/kdb-sidecar/internal/types"
)

type Backup struct {
	base.Api
}

// GetPage
// @Summary 获取集群节点拓扑列表
// @Description 获取集群节点拓扑列表
// @Tags GetBackupPlanList
// @ID get-backup-plan-list
// @Param status query string false "status"
// @Success 200 {object} response.DefaultResponse "{"code": 200, "data": [...]}"
// @Router /api/v1/backup-plan [get]
func (e Backup) GetPage(c *gin.Context) {
	req := types.BackupPlanReq{}
	err := e.MakeContext(c).
		Bind(&req).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	list, err := biz.ListBackupPlan(req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(list, "查询成功")
}

// Modify
// @Summary 修改备份计划
// @Description 修改备份计划
// @Router /api/v1/backup-plan [post]
func (e Backup) Modify(c *gin.Context) {
	req := types.ModifyBackupPlanReq{}
	err := e.MakeContext(c).
		Bind(&req).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	err = biz.ModifyBackupPlan(req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(nil, "查询成功")
}

// Delete
// @Summary 删除备份计划
// @Description 删除备份计划
// @Router /api/v1/backup-plan [delete]
func (e Backup) Delete(c *gin.Context) {
	req := types.DelBackupPlanReq{}
	err := e.MakeContext(c).
		Bind(&req).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	err = biz.DelBackupPlan(req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(nil, "查询成功")
}
