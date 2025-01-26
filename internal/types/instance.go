package types

type InstancesReq struct {
	Status string `json:"status" form:"status" query:"status" search:"type:exact;column:status;table:database_instance"`
}
