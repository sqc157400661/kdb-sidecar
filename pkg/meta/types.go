package meta

type Instance struct {
	ID        string `gorm:"type:varchar(100)" json:"id"`
	ClusterID string `gorm:"type:varchar(100)" json:"clusterID"`
	MasterID  string `gorm:"type:varchar(100)" json:"masterID"`
	ServerID  int    `gorm:"type:int(100)" json:"serverID"`
	PodName   string `gorm:"type:varchar(100)" json:"podName"`
	PodIP     string `gorm:"type:varchar(100)" json:"podIP"`
	Host      string `gorm:"type:varchar(200)" json:"host"`
	Port      int    `gorm:"type:int(100)" json:"port"`
	Role      string `gorm:"type:varchar(20)" json:"role"`
	Version   string `gorm:"type:varchar(20)" json:"version"`
	ReadOnly  bool   `gorm:"type:bool" json:"readOnly"`
	Status    string `gorm:"type:varchar(50)" json:"status"`
	Extra     string `gorm:"type:varchar(255)" json:"extra"`
}
