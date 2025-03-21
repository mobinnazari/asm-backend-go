package config

import "git.sindadsec.ir/asm/backend/utils"

type Config struct {
	Addr          string
	ApiUrl        string
	DisposableUrl string
}

func Init() *Config {
	cfg := &Config{
		Addr:          utils.GetStringEnv("SERVER_ADDR"),
		ApiUrl:        utils.GetStringEnv("EXTERNAL_ADDR"),
		DisposableUrl: utils.GetStringEnv("DISPOSABLE_ADDR"),
	}

	return cfg
}
