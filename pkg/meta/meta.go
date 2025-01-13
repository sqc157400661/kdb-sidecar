package meta

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Manager struct {
	Db *gorm.DB
}

func NewMetaManager() (svc *Manager, err error) {
	// 连接到 SQLite 数据库
	db, err := gorm.Open(sqlite.Open("metadata.db"), &gorm.Config{})
	if err != nil {
		return
	}
	svc = &Manager{Db: db}

	return
}

func (s *Manager) Setup() error {
	// 自动迁移 Metadata 和 AnotherStruct 对应的表
	return s.Db.AutoMigrate(&Instance{})
}
