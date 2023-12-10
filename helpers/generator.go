package helpers

import (
	"math/rand"
	"time"
)

func GenerateRandomID() int {
	rand.Seed(time.Now().UnixNano())
	min := 1000000
	max := 9999999
	return int(min + rand.Intn(max-min+1))
}
