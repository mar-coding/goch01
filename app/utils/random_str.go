package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"
	"unicode"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func initSeed() (seededRand *rand.Rand) {
	// create a seed for random algorithm
	seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	return
}

func GetRandomString(str_len int) (output string) {
	// get a seed and then use it to create random string
	// with length equal to 'str_len'
	if str_len <= 0 {
		str_len = 1
	}
	b := make([]byte, str_len)
	for i := range b {
		b[i] = letters[initSeed().Intn(len(letters))]
	}
	output = string(b)
	return
}

func GetHashStr(str string) string {
	// get a string then hash it with sha256 algorithms
	hash := sha256.Sum256([]byte(str))
	output := hex.EncodeToString(hash[:])
	return output
}

func GetSumOfStr(str string) (sum int) {
	// get a string then split it to its character
	// then check if the char is digit and then use
	// ASCII code to find its value
	sum = 0
	temp := 0
	for _, value := range str {
		if unicode.IsDigit(value) {
			temp = int(value - '0')
			sum += temp
		}
	}
	return
}
