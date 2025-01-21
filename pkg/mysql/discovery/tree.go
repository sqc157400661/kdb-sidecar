package discovery

import "sync"

type InstanceNode struct {
	Host       string
	Port       int
	IsReadonly bool
	IsMaster   bool
	IsBad      bool
	Dept       int
	// ShadowMaster Hot backup node, only used in the case of dual primary nodes
	ShadowMaster *InstanceNode
	Master       *InstanceNode
	Slaves       []*InstanceNode
}

func (in *InstanceNode) Equal(out *InstanceNode) bool {
	if in.Host == out.Host && in.Port == out.Port {
		return true
	}
	return false
}

// Tree 管理树形结构
type Tree struct {
	root  *InstanceNode
	nodes map[string]*InstanceNode // 存储所有节点的映射
	mu    sync.RWMutex             // 保护节点操作的读写锁
}

// NewTree 创建一个新的树实例
func NewTree(root *InstanceNode) *Tree {
	if root == nil {
		root = &InstanceNode{}
	}
	return &Tree{
		root:  root,
		nodes: map[string]*InstanceNode{root.Host: root},
	}
}

func (t *Tree) Exist(node *InstanceNode) bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	_, exists := t.nodes[node.Host]
	return exists
}

func (t *Tree) addNodeIfNotExist(node *InstanceNode) {
	if node == nil {
		return
	}
	t.mu.Lock()
	defer t.mu.Unlock()
	if !t.Exist(node) {
		t.nodes[node.Host] = node
	}
}

func (t *Tree) AddChildNode(parent, child *InstanceNode) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if child == nil {
		return
	}
	if t.Exist(child) && t.Exist(parent) {
		// master <-> shadow master
		if parent.Master.Equal(child) {
			dept := parent.Dept
			if child.Dept != 0 && child.Dept < dept {
				dept = child.Dept
			}
			parent.Dept = dept
			child.Dept = dept
			if parent.IsReadonly {
				parent.Master = child
				child.Master = nil
				child.ShadowMaster = parent
			} else if child.IsReadonly {
				parent.Master = nil
				parent.ShadowMaster = child
				child.Master = parent
			} else {
				parent.ShadowMaster = child
				child.ShadowMaster = parent
				parent.IsBad = true
				child.IsBad = true
			}
		}
		return
	}
	t.addNodeIfNotExist(parent)
	t.addNodeIfNotExist(child)
	if parent == nil {
		t.root = child
	} else {
		child.Master = parent
		child.Dept = parent.Dept + 1
		parent.IsMaster = true
		parent.Slaves = append(parent.Slaves, child)
	}
}
