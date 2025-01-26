package meta

type Instance struct {
	ID         string `xorm:"varchar(100) 'id'" json:"id" tab:"ID"`
	ClusterID  string `xorm:"varchar(100) 'cluster_id'" json:"clusterID" tab:"clusterID"`
	MasterID   string `xorm:"varchar(100) 'master_id'" json:"masterID" tab:"masterID"`
	ServerUUID string `xorm:"varchar(100) 'server_uuid'" json:"serverUUID" tab:"serverUUID"`
	ServerID   int    `xorm:"int 'server_id'" json:"serverID" tab:"serverID"`
	PodName    string `xorm:"varchar(100) 'pod_name'" json:"podName" tab:"podName"`
	PodIP      string `xorm:"varchar(100) 'pod_ip'" json:"podIP" tab:"podIP"`
	Host       string `xorm:"varchar(200) 'host'" json:"host" tab:"host"`
	Port       int    `xorm:"int 'port'" json:"port" tab:"port"`
	Role       string `xorm:"varchar(20) 'role'" json:"role" tab:"role"`
	Version    string `xorm:"varchar(20) 'version'" json:"version" tab:"version"`
	ReadOnly   bool   `xorm:"bool 'read_only'" json:"readOnly" tab:"readOnly"`
	Status     string `xorm:"varchar(50) 'status'" json:"status" tab:"status"`
	Extra      string `xorm:"varchar(255) 'extra'" json:"extra" tab:"extra"`
}
