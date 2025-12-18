package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("===Base 64 Converter by Andini===")
	fmt.Println("Choose an option:")
	fmt.Println("1. Encode")
	fmt.Println("2. Decode")
	fmt.Print("Enter your choice (1 or 2): ")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		encode(reader)
	case "2":
		decode(reader)
	default:
		fmt.Println("Invalid choice. Please run the program again and select 1 or 2.")
	}
}

func encode(reader *bufio.Reader) {
	fmt.Print("\nEnter plain text to encode: ")
	plainText, _ := reader.ReadString('\n')
	plainText = strings.TrimSpace(plainText)

	if plainText == "" {
		fmt.Println("Error: Empty input")
		return
	}

	encoded := base64.StdEncoding.EncodeToString([]byte(plainText))
	fmt.Println("\nEncoded Base64:")
	fmt.Println(encoded)
}

func decode(reader *bufio.Reader) {
	fmt.Print("\nEnter Base64 text to decode: ")
	base64Text, _ := reader.ReadString('\n')
	base64Text = strings.TrimSpace(base64Text)

	if base64Text == "" {
		fmt.Println("Error: Empty input")
		return
	}

	decoded, err := base64.StdEncoding.DecodeString(base64Text)
	if err != nil {
		fmt.Printf("Error decoding Base64: %v\n", err)
		fmt.Println("Make sure the input is valid Base64 text.")
		return
	}

	fmt.Println("\nDecoded plain text:")
	fmt.Println(string(decoded))
}
