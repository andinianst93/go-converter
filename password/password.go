// Package password provides utilities for generating strong, random passwords.
// It ensures passwords meet security best practices with minimum length requirements
// and character diversity (uppercase, lowercase, numbers, and symbols).
package password

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Generate creates 10 strong random passwords with default length of 20 characters
// A strong password has at least 16 characters with mixed-case letters, numbers and symbols
func Generate() {
	fmt.Printf("\n✓ Generated 10 Random Passwords (length: %d each):\n", defaultLength)
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	for i := 1; i <= 10; i++ {
		password := GeneratePassword(defaultLength)
		fmt.Printf("  %2d. %s\n", i, password)
	}

	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("\n✓ These passwords are STRONG because they:")
	fmt.Println("  1. Are at least 16 characters long")
	fmt.Println("  2. Contain mixed-case letters, numbers and symbols")
	fmt.Println("  3. Are completely random")
}

// GenerateWithCustomLength creates a strong random password with user-specified length
// Minimum length is 16 characters to ensure password strength
func GenerateWithCustomLength() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\nEnter desired password length (minimum 16): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	length, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("✗ Invalid input. Please enter a valid number.")
		return
	}

	if length < minLength {
		fmt.Printf("✗ Password length must be at least %d characters for security.\n", minLength)
		fmt.Printf("  Generating with minimum length of %d instead...\n", minLength)
		length = minLength
	}

	for i := 0; i <= 10; i++ {
		password := GeneratePassword(length)
		fmt.Printf("  %s\n", password)
	}
}
