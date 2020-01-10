package textutil

import "math/rand"

const symbols = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// GenerateRandString generates an alphanumeric string of length n to be used as the url identifier
func GenerateRandString(n int) string {
	// three-letter strings allow for arond 230k combinations, ten-letter strings - around 830 quadrillion
	allocatedString := make([]byte, n)
	for i := range allocatedString {
		allocatedString[i] = symbols[rand.Intn(len(symbols))]
	}
	return string(allocatedString)
}
