package util

import (
	"math/rand"
	"strings"
	"time"
)

const (
	alphas = "abcdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphas)

	for i := 0; i < n; i++ {
		c := alphas[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"PLN", "USD", "EUR"}
	return currencies[rand.Intn(len(currencies))]
}
