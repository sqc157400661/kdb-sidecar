package config

const (
	CreateUser      = "CREATE USER %s IF NOT EXISTS"
	GrantGlobalUser = "GRANT %s ON *.* TO '%s'@%s IDENTIFIED BY '%s'"
)
