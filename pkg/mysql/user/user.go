package user

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/sqc157400661/helper/mysql"
	"github.com/sqc157400661/kdb-sidecar/pkg/mysql/config"
	"strings"
)

type Handler struct {
	engine   *xorm.Engine
	executor *mysql.Executor
	users    []config.MySQLUser
}

func NewUserHandle(engine *xorm.Engine, users []config.MySQLUser) *Handler {
	return &Handler{
		engine:   engine,
		users:    users,
		executor: mysql.NewExecutorByEngine(engine),
	}
}

func (c *Handler) Do() error {
	err := c.executor.MakeBinLogOffBySession()
	if err != nil {
		return err
	}
	for _, user := range c.users {
		sql := fmt.Sprintf(config.GrantGlobalUser, strings.Join(user.Privileges, ","), user.Username, user.Host, user.Password)
		_, err = c.engine.SQL(sql).Exec()
		if err != nil {
			return err
		}
	}
	err = c.executor.MakeBinLogOnBySession()
	if err != nil {
		return err
	}
	return nil
}
