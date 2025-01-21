package discovery

import (
	"sync"

	"github.com/go-xorm/xorm"
	"github.com/sqc157400661/helper/mysql"
)

type Discovery interface {
	Discover(host string) error
}

type DiscoverManager struct {
	ReplUser     string `yaml:"repl_user" json:"repl_user"`
	ReplPassword string `yaml:"repl_password" json:"repl_password"`
	currentHost  string
	currentPort  int
	executors    map[string]*mysql.Executor
	sync.RWMutex
}

func (d *DiscoverManager) getRootInfo() (rootNode *InstanceNode, err error) {
	currentExecutor, err := d.getExecutor(d.currentHost, d.currentPort)
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
		masterExecutor, err := d.getExecutor(rootNode.Host, rootNode.Port)
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
				nodeTree.AddChildNode(node, &InstanceNode{Host: row.Host, Port: row.Port})
			}
			if len(slaveHosts) == 0 {
				results, err := parentExecutor.ShowProcesslist()
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
			nextQueue = append(nextQueue, node.Slaves...)
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
		User:   d.ReplUser,
		Passwd: d.ReplPassword,
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

// 获取当前节点的master
// 获取当前节点的所有slaves
// mysql连接到相应的host补全信息
//
