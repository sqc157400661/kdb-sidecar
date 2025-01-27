package config

import (
	"os"
	"strconv"
)

/**
从环境变量中获取相关配置信息
*/

const (
	MySQLPortEnv              = "KDB_PORT"
	MySQLLocalRootEnv         = "MYSQL_LOCAL_ROOT"
	MySQLLocalRootPasswordEnv = "MYSQL_LOCAL_ROOT_PASSWORD"
)

const ClusterIDEnv = "CLUSTER_ID"
const MySQLServerIDEnv = "MYSQL_SERVER_ID"
const K8SNamespaceEnv = "NAMESPACE"
const InitMySQLRoleEnv = "ROLE"
const HostnameEnv = "KDB_HOSTNAME"
const DeployArchEnv = "DEPLOY_ARCH"
const PodIPEnv = "KDB_POD_IP"

var ClusterID string
var MySQLServerID string
var MySQLPort int
var K8SNamespace string
var InitMySQLRole string
var Hostname string
var DeployArch string
var PodIP string

func init() {
	ClusterID = os.Getenv(ClusterIDEnv)
	MySQLServerID = os.Getenv(MySQLLocalRootEnv)
	portStr := os.Getenv(MySQLPortEnv)
	MySQLPort, _ = strconv.Atoi(portStr)
	K8SNamespace = os.Getenv(K8SNamespaceEnv)
	InitMySQLRole = os.Getenv(InitMySQLRoleEnv)
	Hostname = os.Getenv(HostnameEnv)
	DeployArch = os.Getenv(DeployArchEnv)
	PodIP = os.Getenv(PodIPEnv)
}

// oss
const (
	DefaultDaysForOSSStorage = 30

	EndpointENV        = "OPS_OSS_ENDPOINT"
	AccessKeyIdENV     = "OPS_OSS_ACCESS_KEY_ID"
	AccessKeySecretENV = "OPS_OSS_ACCESS_KEY_SECRET"
	BucketENV          = "OPS_OSS_BUCKET"
)
