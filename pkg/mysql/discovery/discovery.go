package discovery

import "github.com/sqc157400661/helper/mysql"

type Discovery interface {
}

type DiscoverManager struct {
	ReplUser     string `yaml:"repl_user" json:"repl_user"`
	ReplPassword string `yaml:"repl_password" json:"repl_password"`
}

func (d *DiscoverManager) GetMasterInfo() {

}

func (d *DiscoverManager) GetSlavesInfo() (slaves []string, err error) {
	executor := mysql.NewExecutorByEngine(nil) // todo add engine
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
