package mysql

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/spf13/cobra"
	"github.com/sqc157400661/helper/mysql"
	"github.com/sqc157400661/kdb-sidecar/internal"
	"github.com/sqc157400661/kdb-sidecar/pkg/mysql/config"
	"github.com/sqc157400661/kdb-sidecar/pkg/mysql/health"
	"github.com/sqc157400661/kdb-sidecar/pkg/mysql/repl"
	"github.com/sqc157400661/kdb-sidecar/pkg/mysql/user"
	"github.com/sqc157400661/kdb-sidecar/pkg/service"
	"github.com/sqc157400661/util"
	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v3"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
	"os"
)

type SidecarOption struct {
	// root mysql user
	RootUser string
	// password of mysql root user
	RootPasswd string
	// the socket file to connect mysql
	RootSocket string
	// mode of service
	Mode int
	// config file for sidecar service
	ConfigFile string
	// configuration parse with configuration file
	config config.MySQLConfig
}

func NewMySQLSidecarServerCmd() *cobra.Command {
	var option = SidecarOption{}
	cmd := &cobra.Command{
		Use:   "MySQLSidecar",
		Short: "Run mysql sidecar server",
		Long:  `Run mysql sidecar server`,
		Run: func(cmd *cobra.Command, args []string) {
			if option.ConfigFile != "" {
				file, err := os.ReadFile(option.ConfigFile)
				if err != nil {
					util.PrintFatalError(err)
				}
				var conf config.MySQLConfig
				err = yaml.Unmarshal(file, &conf)
				if err != nil {
					util.PrintFatalError(err)
				}
				if option.RootUser != "" {
					conf.RootUser = option.RootUser
				}
				if option.RootPasswd != "" {
					conf.RootPasswd = option.RootPasswd
				}
				if option.RootSocket != "" {
					conf.RootSocket = option.RootSocket
				}
				option.config = conf
			}
			err := option.run(args)
			if err != nil {
				klog.Error(err)
				klog.Flush()
				util.PrintFatalError(err)
			}
		},
	}
	cmd.Flags().StringVarP(&option.ConfigFile, "config", "c", "", "Set config file for sidecar service")
	cmd.Flags().StringVarP(&option.RootUser, "user", "u", os.Getenv(config.MySQLLocalRootEnv), "Set root user of mysql for sidecar service")
	cmd.Flags().StringVarP(&option.RootPasswd, "passwd", "p", os.Getenv(config.MySQLLocalRootPasswordEnv), "Set root user password of mysql for sidecar service")
	cmd.Flags().StringVarP(&option.RootSocket, "socket", "s", "/kdbdata/socket/mysqld.sock", "Set socket file to connect mysql")
	cmd.Flags().IntVarP(&option.Mode, "mode", "m", 1, "Set mode of sidecar service,1 mean normal mode，2 mean panic mode")
	return cmd
}

func (o *SidecarOption) run(args []string) (err error) {
	// check if the root user can make a connection to the local database
	var engine *xorm.Engine
	engine, err = mysql.NewMySQLEngine(mysql.ConnectInfo{
		User:   o.config.RootUser,
		Passwd: o.config.RootPasswd,
		Host:   "127.0.0.1",
		Port:   3306,
		Socket: o.config.RootSocket,
	}, true, false)
	if err != nil {
		return err
	}
	// create the MySQL global account that needs to be initialized in the configuration
	err = user.NewUserHandle(engine, o.config.InitUsers).Do()
	if err != nil {
		return err
	}
	// modify mysql.cnf config file
	err = o.modifyMySQLCNFByRole()
	if err != nil {
		return err
	}
	// detecting and building master-slave service
	var cliSet *kubernetes.Clientset
	cliSet, err = internal.GetClientSet()
	if err != nil {
		return err
	}
	services := InitCommonService(engine)
	seeker := repl.NewKubeSeeker(cliSet)
	switch config.DeployArch {
	case internal.MySQLMasterReplicaDeployArch, internal.MySQLMasterSlaveDeployArch:
		services = append(services, repl.NewReplicationService(engine, o.config.Replication, seeker))
	case internal.MySQLSingleDeployArch:
	case internal.MySQLMGRDeployArch:
	default:
		err = fmt.Errorf("not deployArch")
		return err
	}
	//services start
	for _, svc := range services {
		err = svc.Start()
		if err != nil {
			return err
		}
	}
	// detecting and building data backup service
	util.ExitSignalHandler(func() {
		// services stop
		for _, svc := range services {
			err = svc.Stop()
			if err != nil {
				klog.Error(err)
			}
		}
	})
	return nil
}

func (o *SidecarOption) modifyMySQLCNFByRole() error {
	// read mysql my.cnf configuration to obtain readonly and other configurations
	iniConf, err := ini.LoadSources(ini.LoadOptions{
		SkipUnrecognizableLines: true,
	}, o.config.MySQLCNFFile)
	if err != nil {
		return err
	}
	// determine the role status of the current database and add readonly configuration to my.cnf based on the role
	if config.InitMySQLRole == internal.MySQLReplicaRole {
		iniConf.Section("mysqld").Key("read_only").SetValue("1")

	} else {
		iniConf.Section("mysqld").DeleteKey("read_only")
	}
	err = iniConf.SaveTo(o.config.MySQLCNFFile)
	return err
}

func InitCommonService(engine *xorm.Engine) (services []service.Service) {
	// start health check service
	services = append(services, health.NewCheckService(engine, 3))
	return
}
