package internal

const (
	MySQLMasterRole  = "master"
	MySQLReplicaRole = "replica"
)

const (
	MySQLSingleDeployArch = "Single"
	// MySQLMasterSlaveDeployArch Simple Master-Slave,Master->Salve
	MySQLMasterSlaveDeployArch = "Master-Slave"
	// MySQLMasterReplicaDeployArch Master-Replica,Master<->Replica
	MySQLMasterReplicaDeployArch = "Master-Replica"
	MySQLMGRDeployArch           = "MGR"
)
