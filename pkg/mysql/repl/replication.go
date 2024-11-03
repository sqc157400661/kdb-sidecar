package repl

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/sqc157400661/helper/mysql"
	"github.com/sqc157400661/kdb-sidecar/pkg/mysql/config"
	"k8s.io/klog/v2"
	"sync"
	"time"
)

type ReplicationService struct {
	loopSecond  int
	replConf    config.Replication
	seeker      Seeker
	engine      *xorm.Engine
	executor    *mysql.Executor
	startOnce   sync.Once
	stopOnce    sync.Once
	stopChan    chan struct{}
	refreshChan chan struct{}
}

func NewReplicationService(engine *xorm.Engine, replication config.Replication, seeker Seeker) *ReplicationService {
	return &ReplicationService{
		engine:      engine,
		seeker:      seeker,
		loopSecond:  5,
		replConf:    replication,
		refreshChan: make(chan struct{}, 0),
		executor:    mysql.NewExecutorByEngine(engine),
	}
}

// Start service of replication
func (r *ReplicationService) Start() {
	r.startOnce.Do(func() {
		go func() {
			ticker := time.NewTicker(time.Duration(r.loopSecond) * time.Second)
			defer ticker.Stop()
			select {
			case <-ticker.C:
				r.run()
			case <-r.refreshChan:
				err := r.BuildMasterSlave()
				if err != nil {
					klog.Errorf("refresh master err:%s", err.Error())
				}
			case <-r.stopChan:
				close(r.refreshChan)
				klog.Info("health check stopped")
				return
			}
		}()
	})
}

func (r *ReplicationService) run() {
	// verify if master-slave setup is required
	if r.replConf.Host == "" {
		klog.Info("master host is empty, skipped replication")
		return
	}
	// verify master-slave status
	ready, err := r.CheckSlaveStatus()
	if ready {
		klog.Info("master-slave status ready")
		return
	}
	if err != nil {
		klog.Errorf("CheckSlaveStatus err:%s", err.Error())
		err = r.autoRecoverMasterSlave()
		if err != nil {
			klog.Errorf("autoRecoverMasterSlave err:%s", err.Error())
		}
	}
	// build master-slave
	err = r.BuildMasterSlave()
	if err != nil {
		klog.Errorf("build master-slave err: %s", err.Error())
	}
}

func (r *ReplicationService) autoRecoverMasterSlave() error {
	hostInfo, err := r.seeker.GetHostInfoByHostname(r.replConf.Hostname)
	if err != nil {
		return err
	}
	if r.replConf.Host != hostInfo.Host {
		r.replConf.Host = hostInfo.Host
		r.refreshChan <- struct{}{}
	}
	return nil
}

func (r *ReplicationService) CheckSlaveStatus() (ready bool, err error) {
	var status mysql.SlaveStatus
	status, err = r.executor.ShowSlaveStatus()
	if err != nil {
		return false, err
	}
	if status.SlaveIORunning == "Yes" && status.SlaveSQLRunning == "Yes" {
		return true, nil
	}
	if status.LastIOError != "" || status.LastSQLError != "" {
		err = fmt.Errorf("LastIOError:%s LastSQLError:%s", status.LastIOError, status.LastSQLError)
		return
	}
	return
}

func (r *ReplicationService) BuildMasterSlave() (err error) {
	return r.executor.ChangeMasterToWithAuto(r.replConf.Host, config.MySQLPort, r.replConf.ReplUser, r.replConf.ReplPassword)
}

// Stop service of replication
func (r *ReplicationService) Stop() {
	r.stopOnce.Do(func() {
		close(r.stopChan)
	})
}

// RefreshMaster when the master changes, refresh  master
func (r *ReplicationService) RefreshMaster(replication config.Replication) {
	r.replConf = replication
	r.refreshChan <- struct{}{}
}
