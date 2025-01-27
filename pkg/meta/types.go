package meta

import (
	"github.com/sqc157400661/kdb-sidecar/internal"
	"github.com/sqc157400661/kdb-sidecar/pkg/mysql/config"
	"github.com/sqc157400661/kdb-sidecar/pkg/mysql/discovery"
)

type Instance struct {
	ID             string `xorm:"varchar(100) 'id'" json:"id" tab:"id"`
	ClusterID      string `xorm:"varchar(100) 'cluster_id'" json:"clusterID" tab:"clusterID"`
	MasterServerID int    `xorm:"int 'master_server_id'" json:"masterServerID" tab:"masterServerID"`
	ServerUUID     string `xorm:"varchar(100) 'server_uuid'" json:"serverUUID" tab:"serverUUID"`
	ServerID       int    `xorm:"int 'server_id'" json:"serverID" tab:"serverID"`
	PodName        string `xorm:"varchar(100) 'pod_name'" json:"podName" tab:"podName"`
	PodIP          string `xorm:"varchar(100) 'pod_ip'" json:"podIP" tab:"podIP"`
	Host           string `xorm:"varchar(200) 'host'" json:"host" tab:"host"`
	Port           int    `xorm:"int 'port'" json:"port" tab:"port"`
	Dept           int    `xorm:"int 'dept'" json:"dept" tab:"dept"`
	Role           string `xorm:"varchar(20) 'role'" json:"role" tab:"role"`
	Version        string `xorm:"varchar(20) 'version'" json:"version" tab:"version"`
	ReadOnly       bool   `xorm:"bool 'read_only'" json:"readOnly" tab:"readOnly"`
	Status         string `xorm:"varchar(50) 'status'" json:"status" tab:"status"`
	Extra          string `xorm:"varchar(255) 'extra'" json:"extra" tab:"extra"`
}

func (i *Instance) Convert(node *discovery.InstanceNode) {
	i.ClusterID = config.ClusterID
	if node.IsMaster {
		i.Role = internal.MySQLMasterRole
	} else {
		i.Role = internal.MySQLReplicaRole
	}
	if node.ShadowMaster != nil {
		i.MasterServerID = node.ShadowMaster.ServerID
	}
	if node.Master != nil {
		i.MasterServerID = node.Master.ServerID
	}
	i.Version = node.Version
	i.ServerUUID = node.ServerUUID
	i.ServerID = node.ServerID
	i.Host = node.Host
	i.Port = node.Port
	i.Dept = node.Dept

}
