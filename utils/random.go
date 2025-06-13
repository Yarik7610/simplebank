package utils

import (
	"math/rand/v2"
	"strings"
)

func RandomInt64(min, max int64) int64 {
	return rand.Int64N(max-min) + min
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for range n {
		ch := alphabet[rand.IntN(k)]
		sb.WriteByte(ch)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt64(-1000, 1000)
}

var currencies = []string{"EUR", "USD", "CAD"}

func RandomCurrency() string {
	l := len(currencies)
	return currencies[rand.IntN(l)]
}
