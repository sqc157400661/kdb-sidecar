package health

import (
	"github.com/go-xorm/xorm"
	"github.com/sqc157400661/helper/mysql"
	"github.com/sqc157400661/util"
	"k8s.io/klog/v2"
	"sync"
	"time"
)

const (
	CreateHealthTableSQL = "CREATE TABLE IF NOT EXISTS `health_check` ( `id` int(11) NOT NULL, `t_modified` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY (`id`)) ENGINE=InnoDB"
	CheckSQL             = "insert into mysql.health_check(id, t_modified) values (@@server_id, now()) on duplicate key update t_modified=now()"
)

type CheckService struct {
	engine     *xorm.Engine
	loopSecond int
	executor   *mysql.Executor
	startOnce  sync.Once
	stopOnce   sync.Once
	stopChan   chan struct{}
}

func NewCheckService(engine *xorm.Engine, loopSecond int) *CheckService {
	return &CheckService{
		engine:     engine,
		loopSecond: loopSecond,
		executor:   mysql.NewExecutorByEngine(engine),
	}
}

// Start  the health check service
func (s *CheckService) Start() {
	s.startOnce.Do(func() {
		err := s.CreateTableIfNotExist()
		if err != nil {
			util.PrintFatalError(err)
		}
		go func() {
			ticker := time.NewTicker(time.Duration(s.loopSecond) * time.Second)
			defer ticker.Stop()
			select {
			case <-ticker.C:
				err = s.doCheck()
				if err != nil {
					klog.Errorf("check health err: %s", err.Error())
				}
			case <-s.stopChan:
				klog.Info("health check stopped")
				return
			}
		}()
	})
}

// CreateTableIfNotExist create health check table if not exist
func (s *CheckService) CreateTableIfNotExist() error {
	err := s.executor.MakeBinLogOffBySession()
	if err != nil {
		return err
	}
	_, err = s.engine.SQL("use mysql").Exec()
	if err != nil {
		return err
	}
	_, err = s.engine.SQL(CreateHealthTableSQL).Exec()
	if err != nil {
		return err
	}
	err = s.executor.MakeBinLogOnBySession()
	return err
}

func (s *CheckService) doCheck() error {
	err := s.executor.MakeBinLogOffBySession()
	if err != nil {
		return err
	}
	_, err = s.engine.SQL(CheckSQL).Exec()
	// TODO 记录Log
	return err
}

// Stop  the health check service
func (s *CheckService) Stop() {
	s.stopOnce.Do(func() {
		close(s.stopChan)
	})
}
