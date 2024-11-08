package config

import (
	"os"
	"strconv"
)

/**
从环境变量中获取相关配置信息
*/

const (
	MySQLPortEnv              = "MySQLPort"
	MySQLLocalRootEnv         = "MySQLLocalRoot"
	MySQLLocalRootPasswordEnv = "MySQLLocalRootPassword"
)

const ClusterIDEnv = "ClusterID"
const MySQLServerIDEnv = "MySQLServerID"
const K8SNamespaceEnv = "K8SNamespace"
const InitMySQLRoleEnv = "InitMySQLRole"

var ClusterID string
var MySQLServerID string
var MySQLPort int
var K8SNamespace string
var InitMySQLRole string

func init() {
	ClusterID = os.Getenv(ClusterIDEnv)
	MySQLServerID = os.Getenv(MySQLLocalRootEnv)
	portStr := os.Getenv(MySQLPortEnv)
	MySQLPort, _ = strconv.Atoi(portStr)
	K8SNamespace = os.Getenv(K8SNamespaceEnv)
	InitMySQLRole = os.Getenv(InitMySQLRoleEnv)
}

// oss
const (
	DefaultDaysForOSSStorage = 30

	EndpointENV        = "OPS_OSS_ENDPOINT"
	AccessKeyIdENV     = "OPS_OSS_ACCESS_KEY_ID"
	AccessKeySecretENV = "OPS_OSS_ACCESS_KEY_SECRET"
	BucketENV          = "OPS_OSS_BUCKET"
)
