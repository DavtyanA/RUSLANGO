package commands

import (
	"math/rand"
	"strconv"
	"time"
)

func Roll(input int) string {
	min := 0
	max := input + 1
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(max-min) + min)
}
