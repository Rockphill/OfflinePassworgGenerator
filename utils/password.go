package utils

import (
	"math/rand"
	"time"
)

// GeneratePassword генерирует пароль заданной длины
func GeneratePassword(length int, includeSpecial bool) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	specials := "!@#$%&*()-_=+;:,.<>?/"

	charset := letters
	if includeSpecial {
		charset += specials
	}

	rand.Seed(time.Now().UnixNano())
	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	return string(password)
}
