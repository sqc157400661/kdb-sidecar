package meta

type Instance struct {
	ID         string `xorm:"varchar(100) 'id'" json:"id"`
	ClusterID  string `xorm:"varchar(100) 'cluster_id'" json:"clusterID"`
	MasterID   string `xorm:"varchar(100) 'master_id'" json:"masterID"`
	ServerUUID string `xorm:"varchar(100) 'server_uuid'" json:"serverUUID"`
	ServerID   int    `xorm:"int 'server_id'" json:"serverID"`
	PodName    string `xorm:"varchar(100) 'pod_name'" json:"podName"`
	PodIP      string `xorm:"varchar(100) 'pod_ip'" json:"podIP"`
	Host       string `xorm:"varchar(200) 'host'" json:"host"`
	Port       int    `xorm:"int 'port'" json:"port"`
	Role       string `xorm:"varchar(20) 'role'" json:"role"`
	Version    string `xorm:"varchar(20) 'version'" json:"version"`
	ReadOnly   bool   `xorm:"bool 'read_only'" json:"readOnly"`
	Status     string `xorm:"varchar(50) 'status'" json:"status"`
	Extra      string `xorm:"varchar(255) 'extra'" json:"extra"`
}
