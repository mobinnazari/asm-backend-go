package config

import "git.sindadsec.ir/asm/backend/utils"

type Config struct {
	Addr          string
	MysqlAddr     string
	ApiUrl        string
	DisposableUrl string
	RedisAddr     string
	Mail          MailConfig
}

type MailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

func Init() *Config {
	cfg := &Config{
		Addr:          utils.GetStringEnv("SERVER_ADDR"),
		MysqlAddr:     utils.GetStringEnv("MYSQL_ADDR"),
		ApiUrl:        utils.GetStringEnv("EXTERNAL_ADDR"),
		DisposableUrl: utils.GetStringEnv("DISPOSABLE_ADDR"),
		RedisAddr:     utils.GetStringEnv("REDIS_ADDR"),
		Mail: MailConfig{
			Host:     utils.GetStringEnv("MAIL_HOST"),
			Port:     utils.GetIntEnv("MAIL_PORT"),
			Username: utils.GetStringEnv("MAIL_USERNAME"),
			Password: utils.GetStringEnv("MAIL_PASSWORD"),
		},
	}

	return cfg
}
