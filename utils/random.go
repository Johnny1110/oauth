package utils

import (
	"math/rand"
	"strings"
	"time"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberBytes   = "0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	numberIdxBits = 4                    // 4 bits to represent a number index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	numberIdxMask = 1<<numberIdxBits - 1 // All 1-bits, as many as numberIdxBits
)

func init() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
}

func GenRandomString(length int, allUpperCase bool) string {
	// Create a builder with the specified capacity
	var sb strings.Builder
	sb.Grow(length)

	for i := 0; i < length; i++ {
		// Randomly select from letters or numbers
		if rand.Intn(2) == 0 {
			// Random letter
			idx := rand.Intn(len(letterBytes))
			char := letterBytes[idx]
			if allUpperCase {
				char = strings.ToUpper(string(char))[0]
			}
			sb.WriteByte(char)
		} else {
			// Random number
			idx := rand.Intn(len(numberBytes))
			sb.WriteByte(numberBytes[idx])
		}
	}

	return sb.String()
}
