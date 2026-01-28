package shortener

import (
	"math/rand/v2"
)

func GenerateShortID() string {
	var letters = "0123456789abcdefghijklmnopqrstuvwxyz"
	var length int = 4

	result := make([]byte, length)
	for i := range length {
		result[i] = letters[rand.IntN(len(letters))]
	}

	return string(result)
}
