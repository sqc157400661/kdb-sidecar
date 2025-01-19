package meta

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

type Manager struct {
	Db *xorm.Engine
}

var manager *Manager

func init() {
	manager, _ = NewMetaManager()
}

func DB() *xorm.Engine {
	return manager.Db
}

func NewMetaManager() (svc *Manager, err error) {
	// 连接到 SQLite 数据库
	engine, err := xorm.NewEngine("sqlite3", "metadata.db")
	if err != nil {
		return
	}
	svc = &Manager{Db: engine}

	return
}

func (s *Manager) Setup() error {
	// 自动迁移 Metadata 和 AnotherStruct 对应的表
	return s.Db.Sync2(&Instance{})
}
