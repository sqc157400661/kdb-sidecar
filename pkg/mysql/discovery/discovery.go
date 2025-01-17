package discovery

import (
	"github.com/go-xorm/xorm"
	"github.com/sqc157400661/helper/mysql"
)

type Discovery interface {
}

type DiscoverManager struct {
	ReplUser     string `yaml:"repl_user" json:"repl_user"`
	ReplPassword string `yaml:"repl_password" json:"repl_password"`
}

func (d *DiscoverManager) GetMasterInfo() {

}

func (d *DiscoverManager) GetSlavesInfo(host string) (slaves []string, err error) {
	// todo get instance info by host?
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
	executor := mysql.NewExecutorByEngine(engine)
	results, err := executor.ShowProcesslist()
	if err != nil {
		return
	}
	for _, row := range results {
		if row.State == "Slave" {
			slaves = append(slaves, row.Host)
		}
	}
	return
}
