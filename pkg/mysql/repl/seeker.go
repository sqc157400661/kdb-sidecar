package repl

type HostInfo struct {
	Host string `yaml:"host" json:"host"`
	IP   string `yaml:"ip" json:"ip"`
	Port int    `yaml:"port" json:"port"`
}
type Seeker interface {
	GetHostInfoByPodName(podName string) (*HostInfo, error)
	GetHostInfoByClusterID(id string) ([]*HostInfo, error)
}
