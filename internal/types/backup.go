package types

type BackupPlanReq struct {
	Status string `json:"status" form:"status" query:"status" search:"type:exact;column:status;table:database_instance"`
}

type ModifyBackupPlanReq struct {
	BackupType string `json:"backup_type,omitempty"` // 备份类型 full incr
	Days       string `json:"days,omitempty"`        // 全量备份周的计划
	Hour       string `json:"hour,omitempty"`        // 全量备份小时的计划
}

type DelBackupPlanReq struct {
	BackupType string `json:"backup_type,omitempty"` // 备份类型 full incr all
}
