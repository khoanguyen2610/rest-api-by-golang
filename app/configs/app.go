package configs

import (
	"github.com/kelseyhightower/envconfig"
)
type AppConfig struct {
	DB MysqlConfig
}

type MysqlConfig struct {
	Host   			string	`envconfig:"MYSQL_HOST" required:"true"`
	Port   			int		`envconfig:"MYSQL_PORT" required:"true"`
	DbName 			string	`envconfig:"MYSQL_DB" required:"true"`
	User   			string	`envconfig:"MYSQL_USER" required:"true"`
	Pass   			string	`envconfig:"MYSQL_PASS" required:"true"`
	DbLogEnable   	bool	`envconfig:"MYSQL_LOG_MODE_ENABLE"`
}

func GetAppConfigFromEnv() AppConfig {
	var conf AppConfig
	envconfig.Process("myapp", &conf)
	return conf
}
