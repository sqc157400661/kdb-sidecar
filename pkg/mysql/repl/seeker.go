package repl

type Seeker interface {
	GetHostInfoByHostname()
	GetHostInfoByClusterID()
}
