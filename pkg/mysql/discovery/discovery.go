package discovery

import (
	"sync"

	"github.com/go-xorm/xorm"
	"github.com/sqc157400661/helper/mysql"
)

type Discovery interface {
	Discover(host string) error
}

type option struct {
	currentHost string
	currentPort int
}
type OptionFunc func(*option)

func WithCurrentHost(host string) OptionFunc {
	return func(o *option) {
		o.currentHost = host
	}
}

func WithCurrentPort(port int) OptionFunc {
	return func(o *option) {
		o.currentPort = port
	}
}

type DiscoverManager struct {
	option
	replUser     string `yaml:"repl_user"`
	replPassword string `yaml:"repl_password"`
	executors    map[string]*mysql.Executor
	sync.RWMutex
}

func NewDiscoveryManager(replUser, replPassword string, opts ...OptionFunc) *DiscoverManager {
	var opt = option{
		currentHost: "localhost",
		currentPort: 3306,
	}
	for _, fn := range opts {
		fn(&opt)
	}
	return &DiscoverManager{
		option:       opt,
		executors:    map[string]*mysql.Executor{},
		replUser:     replUser,
		replPassword: replPassword,
	}
}

func (d *DiscoverManager) getRootInfo() (rootNode *InstanceNode, err error) {
	currentExecutor, err := d.getExecutor(d.currentHost, d.currentPort)
	if err != nil {
		return
	}
	var slaveStatus mysql.SlaveStatus
	var masterExecutor *mysql.Executor
	slaveStatus, err = currentExecutor.ShowSlaveStatus()
	if err != nil {
		return
	}
	rootNode = &InstanceNode{}
	for slaveStatus.MasterHost != "" {
		rootNode.Host = slaveStatus.MasterHost
		rootNode.Port = slaveStatus.MasterPort
		masterExecutor, err = d.getExecutor(rootNode.Host, rootNode.Port)
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

func (d *DiscoverManager) findSlavesInfo(rootNode *InstanceNode) (nodeTree *Tree, err error) {
	// todo 根据rootNode循环查找其从实例
	var parentExecutor *mysql.Executor
	var slaveHosts []*mysql.SlaveHost
	nodeTree = NewTree(rootNode)
	var queue = []*InstanceNode{rootNode}
	dept := 1
	for len(queue) > 0 {
		var nextQueue []*InstanceNode
		for _, node := range queue {
			parentExecutor, err = d.getExecutor(node.Host, node.Port)
			if err != nil {
				return
			}
			node.Version, err = parentExecutor.Version()
			if err != nil {
				return
			}
			node.ServerID, err = parentExecutor.ServerID()
			if err != nil {
				return
			}
			// Attempt SHOW SLAVE HOSTS before PROCESSLIST
			slaveHosts, err = parentExecutor.ShowSlaveHosts()
			if err != nil {
				return
			}
			node.IsReadonly, err = parentExecutor.IsReadOnly()
			if err != nil {
				return
			}
			for _, row := range slaveHosts {
				nodeTree.AddChildNode(node, &InstanceNode{
					Host:       row.Host,
					Port:       row.Port,
					ServerUUID: row.SlaveUUID,
					ServerID:   row.ServerID,
				})
			}
			if len(slaveHosts) == 0 {
				var results []mysql.ProcessList
				results, err = parentExecutor.ShowProcesslist()
				if err != nil {
					return
				}
				for _, row := range results {
					if row.State == "Slave" {
						// todo: need confirm port when use show process list
						nodeTree.AddChildNode(node, &InstanceNode{Host: row.Host, Port: 3306})
					}
				}
			}
			if len(node.Slaves) > 0 {
				nextQueue = append(nextQueue, node.Slaves...)
			}
		}
		queue = nextQueue
		dept++
	}
	return
}

func (d *DiscoverManager) getExecutor(host string, port int) (executor *mysql.Executor, err error) {
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
		User:   d.replUser,
		Passwd: d.replPassword,
		Host:   host,
		Port:   port,
	}, true, false)
	if err != nil {
		return
	}
	executor = mysql.NewExecutorByEngine(engine)
	d.executors[host] = executor

	return
}

func (d *DiscoverManager) Discover() (nodeTree *Tree, err error) {
	var rootNode *InstanceNode
	rootNode, err = d.getRootInfo()
	if err != nil {
		return
	}
	nodeTree, err = d.findSlavesInfo(rootNode)
	return
}
