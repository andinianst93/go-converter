package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/andinianst94/go-converter/base64"
	"github.com/andinianst94/go-converter/password"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	// User must choose if they want to generate a password or encode/decode Base64 text.
	fmt.Println("===Welcome to DevTools===")
	fmt.Println("Choose an option:")
	fmt.Println("1. Password Generator")
	fmt.Println("2. Base64 Converter")
	fmt.Print("Enter your choice (1 or 2): ")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		// Password Generator
		fmt.Println("\n===Password Generator===")
		fmt.Println("Choose an option:")
		fmt.Println("1. Generate Password")
		fmt.Println("2. Generate Password with Custom Length")
		fmt.Print("Enter your choice (1 or 2): ")

		subChoice, _ := reader.ReadString('\n')
		subChoice = strings.TrimSpace(subChoice)

		switch subChoice {
		case "1":
			password.Generate()
		case "2":
			password.GenerateWithCustomLength()
		default:
			fmt.Println("Invalid choice. Please run the program again and select 1 or 2.")
		}

	case "2":
		// Base64
		fmt.Println("\n===Base 64 Converter===")
		fmt.Println("Choose an option:")
		fmt.Println("1. Encode")
		fmt.Println("2. Decode")
		fmt.Print("Enter your choice (1 or 2): ")

		subChoice, _ := reader.ReadString('\n')
		subChoice = strings.TrimSpace(subChoice)

		switch subChoice {
		case "1":
			base64.Encode(reader)
		case "2":
			base64.Decode(reader)
		default:
			fmt.Println("Invalid choice. Please run the program again and select 1 or 2.")
		}

	default:
		fmt.Println("Invalid choice. Please run the program again and select 1 or 2.")
	}
}
