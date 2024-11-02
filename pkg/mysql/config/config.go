package config

type MySQLUser struct {
	Username   string   `yaml:"crontab" json:"username"`
	Password   string   `yaml:"crontab" json:"password"`
	Host       string   `yaml:"crontab" json:"host"`
	Privileges []string `yaml:"crontab" json:"privileges"`
}

type Replication struct {
	Host         string `yaml:"host" json:"host"`
	ReplUser     string `yaml:"crontab" json:"repl_user"`
	ReplPassword string `yaml:"crontab" json:"repl_password"`
}

type OssConfig struct{}
type S3Config struct{}
type Backup struct {
	Crontab string    `yaml:"crontab" json:"crontab"`
	Oss     OssConfig `yaml:"crontab" json:"oss"`
	S3      S3Config  `yaml:"crontab" json:"s3"`
}
type MySQLConfig struct {
	// root mysql user
	RootUser string `yaml:"root_user" json:"root_user"`
	// password of mysql root user
	RootPasswd string `yaml:"root_password" json:"root_password"`
	// the socket file to connect mysql
	RootSocket string `yaml:"root_socket" json:"root_socket"`
	// the file of mysql cnf configuration
	MySQLCNFFile string `yaml:"mysql_cnf_file" json:"mysql_cnf_file"`
	// Initialize  global users
	InitUsers []MySQLUser `json:"init_users"`
	// the config info of replication
	Replication Replication `json:"replication"`
	// the config info of backup
	Backup Backup `json:"backup"`
}
