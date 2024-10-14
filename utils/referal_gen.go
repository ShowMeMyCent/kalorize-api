package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func GenerateReferalCode(fullname string) string {
	username := strings.Split(fullname, " ")[0]
	return username + RandomInt(5)
}

func RandomInt(n int) string {
	rand.Seed(time.Now().UnixNano())

	min := int64(1)
	max := int64(10)

	for i := 1; i < n; i++ {
		max *= 10
	}

	randomInt := rand.Int63n(max-min) + min
	return strconv.FormatInt(randomInt, 10)

}
