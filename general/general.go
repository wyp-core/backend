package general

import (
	"math/rand"
	"time"
)

// InArr checks for value present in a slice / array
func InArr[T string | int | int16 | int32 | int64 | bool](val T, arr []T) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func GenerateRandom4DigitSecure() (int, error) {
	// Generate random number between 1000 and 9999
	rand.Seed(time.Now().UnixNano()) // seed once
	num := rand.Intn(9000) + 1000    // range: 1000–9999
	return num, nil
}
