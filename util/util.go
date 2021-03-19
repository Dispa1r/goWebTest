package util

import (
	"math/rand"
	"time"
)

func RandomString(n int) string{
	var letters = []byte("asdfghjklqwertyuiopzxcvbnm1234567890")
	result := make([]byte,n)
	rand.Seed(time.Now().Unix())
	for i:= range result{
		result[i]=letters[rand.Intn(len(letters))]
	}
	return string(result)
}