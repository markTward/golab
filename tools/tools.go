package tools

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	var letters_numbers = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rs := make([]rune, n)

	// seed random generator
	rand.Seed(time.Now().UTC().UnixNano())

	// randomly fill rune array
	for i := range rs {
		rs[i] = letters_numbers[rand.Intn(len(letters_numbers))]
	}
	return string(rs)
}
