package configs

type AppConfig struct {
	DB MysqlConfig
}

type MysqlConfig struct {
	Host   string
	Port   int
	DbName string
	User   string
	Pass   string
}

func GetAppConfigFromEnv() AppConfig {
	conf := AppConfig{
		DB: MysqlConfig{
			Host:   "mysql-server",
			Port:   3306,
			DbName: "sample_go",
			User:   "root",
			Pass:   "123456",
		},
	}
	return conf
}
