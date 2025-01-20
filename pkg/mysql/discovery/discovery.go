package discovery

import (
	"github.com/go-xorm/xorm"
	"github.com/sqc157400661/helper/mysql"
	"sync"
)

type Discovery interface {
	Discover(host string) error
}

type DiscoverManager struct {
	ReplUser     string `yaml:"repl_user" json:"repl_user"`
	ReplPassword string `yaml:"repl_password" json:"repl_password"`
	currentHost  string
	executors    map[string]*mysql.Executor
	sync.RWMutex
}

func (d *DiscoverManager) resolvedMasterInfo() (rootNode *InstanceNode, err error) {
	currentExecutor, err := d.getExecutor(d.currentHost)
	if err != nil {
		return
	}
	var slaveStatus mysql.SlaveStatus
	slaveStatus, err = currentExecutor.ShowSlaveStatus()
	if err != nil {
		return
	}
	rootNode = &InstanceNode{}
	for slaveStatus.MasterHost != "" {
		rootNode.Host = slaveStatus.MasterHost
		rootNode.Port = slaveStatus.MasterPort
		masterExecutor, err := d.getExecutor(rootNode.Host)
		if err != nil {
			return
		}
		slaveStatus, err = masterExecutor.ShowSlaveStatus()
		if err != nil {
			return
		}
		if slaveStatus.MasterHost == d.currentHost {
			break
		}
	}
	return
}

func (d *DiscoverManager) findSlavesInfo(host string) (slaves []string, err error) {
	// todo 根据rootNode循环查找其从实例
	var executor *mysql.Executor
	executor, err = d.getExecutor(host)
	if err != nil {
		return
	}
	var slaveHosts []*mysql.SlaveHost
	// Attempt SHOW SLAVE HOSTS before PROCESSLIST
	slaveHosts, err = executor.ShowSlaveHosts()
	if err != nil {
		return
	}
	for _, row := range slaveHosts {
		slaves = append(slaves, row.Host)
	}
	if len(slaveHosts) == 0 {
		results, err := executor.ShowProcesslist()
		if err != nil {
			return
		}
		for _, row := range results {
			if row.State == "Slave" {
				slaves = append(slaves, row.Host)
			}
		}
	}

	return
}

func (d *DiscoverManager) getExecutor(host string) (executor *mysql.Executor, err error) {
	d.RLock()
	existExecutor, exist := d.executors[host]
	d.RUnlock()
	if exist {
		return existExecutor, nil
	}
	d.Lock()
	defer d.Unlock()
	var engine *xorm.Engine
	engine, err = mysql.NewMySQLEngine(mysql.ConnectInfo{
		User:   d.ReplUser,
		Passwd: d.ReplPassword,
		Host:   host,
		Port:   3306,
	}, true, false)
	if err != nil {
		return
	}
	executor = mysql.NewExecutorByEngine(engine)
	d.executors[host] = executor

	return
}

func (d *DiscoverManager) Discover() {

}

// 获取当前节点的master
// 获取当前节点的所有slaves
// mysql连接到相应的host补全信息
//
