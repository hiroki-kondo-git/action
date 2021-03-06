package util

import (
	"errors"
	"math/rand"
	"time"
)

func RandomString(n uint32) (string, error) {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	// lettersの〜番目（数字）を10こ生成
	numberOrder := make([]byte, n)
	if _, err := rand.Read(numberOrder); err != nil {
		return "", errors.New("unexpected error")
	}

	var result string
	for _, v := range numberOrder {
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}

func RandomInt(n int) int {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(n)
	return num
}
