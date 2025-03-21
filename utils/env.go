package utils

import (
	"log"
	"os"
	"strconv"
)

func GetStringEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Panicf("Failed to load env %s", key)
	}

	return val
}

func GetIntEnv(key string) int {
	val := GetStringEnv(key)

	intVal, err := strconv.Atoi(val)
	if err != nil {
		log.Panicf("Failed to convert env %s to integer", key)
	}

	return intVal
}
