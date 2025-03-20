package utils

import (
	"log"
	"os"
)

func GetStringEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Panicf("Failed to load env %s", key)
	}

	return val
}
