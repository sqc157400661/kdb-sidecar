package meta_service

import (
	"github.com/sqc157400661/helper/mysql"
	"github.com/sqc157400661/kdb-sidecar/pkg/meta"
	"github.com/sqc157400661/kdb-sidecar/pkg/mysql/config"
	"github.com/sqc157400661/kdb-sidecar/pkg/mysql/discovery"
	"k8s.io/klog/v2"
	"sync"
	"time"
)

type MetaService struct {
	executor   *mysql.Executor
	config     config.MySQLConfig
	loopSecond int
	startOnce  sync.Once
	stopOnce   sync.Once
	stopChan   chan struct{}
}

func NewMetaService(executor *mysql.Executor, config config.MySQLConfig) *MetaService {
	return &MetaService{
		executor:   executor,
		config:     config,
		loopSecond: 60,
	}
}

func (s *MetaService) Start() error {
	var err error
	s.startOnce.Do(func() {
		go func() {
			ticker := time.NewTicker(time.Duration(s.loopSecond) * time.Second)
			defer ticker.Stop()
			select {
			case <-ticker.C:
				err = s.sync()
				if err != nil {
					klog.Errorf("sync meta err: %s", err.Error())
				}
			case <-s.stopChan:
				klog.Info("sync meta service stopped")
				return
			}
		}()
	})
	return err
}

func (s *MetaService) sync() (err error) {
	// TODO: raft master do discover
	var tree *discovery.Tree
	discover := discovery.NewDiscoveryManager(s.config.Replication.ReplUser, s.config.Replication.ReplPassword)
	tree, err = discover.Discover()
	if err != nil {
		return err
	}
	tree.ForEach(func(node *discovery.InstanceNode) {
		if node == nil {
			instance := &meta.Instance{}
			instance.Convert(node)
			if instance.ServerID > 0 {
				// TODO 尝试通过接口查询podip和podname
				_, err = meta.DB().Insert(instance)
				if err != nil {
					// todo add errors log
				}
			}
		}
	})
	return err
}

func (s *MetaService) Stop() error {
	s.stopOnce.Do(func() {
		close(s.stopChan)
	})
	return nil
}
