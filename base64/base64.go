package base64

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"strings"
)

func Encode(reader *bufio.Reader) {
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

func Decode(reader *bufio.Reader) {
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
