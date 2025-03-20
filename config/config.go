package config

import "git.sindadsec.ir/asm/backend/utils"

type Config struct {
	Addr string
}

func Init() *Config {
	cfg := &Config{
		Addr: utils.GetStringEnv("SERVER_ADDR"),
	}

	return cfg
}
