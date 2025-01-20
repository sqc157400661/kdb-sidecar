package discovery

import "sync"

type InstanceNode struct {
	Host     string
	Port     int
	IsMaster bool
	Dept     int
	Master   *InstanceNode
	Slaves   []*InstanceNode
}

// Tree 管理树形结构
type Tree struct {
	root  *InstanceNode
	nodes map[string]*InstanceNode // 存储所有节点的映射
	mu    sync.RWMutex             // 保护节点操作的读写锁
}

// NewTree 创建一个新的树实例
func NewTree(rootHost string) *Tree {
	root := &InstanceNode{Host: rootHost}
	return &Tree{
		root:  root,
		nodes: map[string]*InstanceNode{rootHost: root},
	}
}

// AddNode 向树中添加一个子节点，指定父节点ID
func (t *Tree) AddNode(parentHost, host string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	// 如果父节点不存在，创建一个新的父节点
	parentNode, exists := t.nodes[parentHost]
	if !exists {
		parentNode = &InstanceNode{Host: parentHost}
		t.nodes[parentHost] = parentNode
	}

	// 创建子节点
	childNode := &InstanceNode{Host: host, Master: parentNode}
	parentNode.Slaves = append(parentNode.Slaves, childNode)

	// 将子节点加入到树中
	t.nodes[host] = childNode
}
