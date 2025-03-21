package utils

import (
	"math/rand"
	"strconv"
)

func GenerateRandomCode() string {
	randInt := rand.Intn(900000) + 100000
	return strconv.Itoa(randInt)
}
