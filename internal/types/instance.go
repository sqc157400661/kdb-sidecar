package types

type InstancesReq struct {
	ClusterName string `json:"cluster_name" form:"cluster_name" query:"cluster_name" search:"type:exact;column:cluster_name;table:database_instance"`
	Status      string `json:"status" form:"status" query:"status" search:"type:exact;column:status;table:database_instance"`
}
