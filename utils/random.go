package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// randon name owner
func RandomOwner() string {
	return RandomString(6)
}

// random money
func RandomMoney() int64 {
	return RandomInt(0, 1_000_000_000)
}

// RandomCurrency
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "VND"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func RandomIdInt() int64 {
	return rand.Int63()
}
