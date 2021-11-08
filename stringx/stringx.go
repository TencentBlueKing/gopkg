package stringx

import "math/rand"

// Truncate string to specific length
func Truncate(s string, n int) string {
	if n > len(s) {
		return s
	}
	return s[:n]
}


const letterBytes = "abcdefghijklmnopqrstuvwxyz1234567890"

// Random generate a random string with fixed length. [a-z0-9]
func Random(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
