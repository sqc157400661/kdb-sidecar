package biz

import (
	"github.com/sqc157400661/kdb-sidecar/internal/types"
	"github.com/sqc157400661/kdb-sidecar/pkg/meta"
)

// ListBackupPlan the list of backup plan
func ListBackupPlan(req types.BackupPlanReq) (plans []*meta.BackupPlan, err error) {
	err = meta.DB().Find(&plans)
	return
}

// ModifyBackupPlan TODO:
func ModifyBackupPlan(req types.ModifyBackupPlanReq) (err error) {
	return
}

// DelBackupPlan TODO:
func DelBackupPlan(req types.DelBackupPlanReq) (err error) {
	return
}
