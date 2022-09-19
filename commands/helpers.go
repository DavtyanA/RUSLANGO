package commands

import (
	"math/rand"
	"strings"
	"time"
)

// Ebanya Xyeta
func GetRandomItem[T interface{}](inputarray []T) T {
	rand.Seed(time.Now().UnixNano())
	//not sure if I need (len - 1), examples don't do that
	randomItem := inputarray[rand.Intn(len(inputarray))]
	return randomItem
}

func StringContains(S string, sub string) bool {
	return strings.Contains(strings.ToLower(S), strings.ToLower(sub))
}

func StringStartsWith(S string, sub string) bool {
	return strings.HasPrefix(strings.ToLower(S), strings.ToLower(sub))
}

func StringContainsArray(S string, subs []string) bool {
	for _, s := range subs {
		if StringContains(S, s) {
			return true
		}
	}
	return false
}
