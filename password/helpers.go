package password

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	digits           = "0123456789"
	symbols          = "!@#$%^&*()_+-=[]{}|;:,.<>?"
	allCharacters    = uppercaseLetters + lowercaseLetters + digits + symbols
	defaultLength    = 20
	minLength        = 16
)

// generatePassword creates a random password with specified length
// It ensures the password contains at least one character from each category:
// uppercase, lowercase, digit, and symbol
func GeneratePassword(length int) string {
	if length < minLength {
		length = minLength
	}

	password := make([]byte, length)

	// Ensure at least one character from each category
	password[0] = uppercaseLetters[randomInt(len(uppercaseLetters))]
	password[1] = lowercaseLetters[randomInt(len(lowercaseLetters))]
	password[2] = digits[randomInt(len(digits))]
	password[3] = symbols[randomInt(len(symbols))]

	// Fill the rest with random characters from all categories
	for i := 4; i < length; i++ {
		password[i] = allCharacters[randomInt(len(allCharacters))]
	}

	// Shuffle the password to randomize position of guaranteed characters
	for i := length - 1; i > 0; i-- {
		j := randomInt(i + 1)
		password[i], password[j] = password[j], password[i]
	}

	return string(password)
}

// randomInt generates a cryptographically secure random integer between 0 and max-1
func randomInt(max int) int {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		// Fallback to a simple method if crypto/rand fails (very unlikely)
		panic(fmt.Sprintf("failed to generate random number: %v", err))
	}
	return int(n.Int64())
}
