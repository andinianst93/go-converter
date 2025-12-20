# DevTools - Password Generator & Base64 Converter

A simple command-line tool written in Go for generating strong passwords and encoding/decoding Base64 text.

## Features

### Password Generator
- **Generate 10 Strong Passwords**: Instantly create 10 random, secure passwords with default 20 characters each
- **Custom Length Support**: Generate 10 passwords with your specified length (minimum 16 characters)
- **Strong Security**: Uses cryptographically secure random generation (`crypto/rand`)
- **Character Diversity**: Includes uppercase, lowercase, numbers, and symbols
- **Best Practices**: Follows industry standards for password strength

### Base64 Converter
- **Encode**: Convert plain text to Base64
- **Decode**: Convert Base64 back to plain text
- **Error Handling**: Validates invalid Base64 inputs
- **Interactive**: Easy-to-use command-line interface

## Password Strength Guidelines

A strong password follows ALL THREE of these tips:

### 1. Make them long
- **Minimum**: At least 16 characters
- **Default**: 20 characters
- **Longer is stronger!**

### 2. Make them random
This tool generates random strings with mixed-case letters, numbers, and symbols. For example:
```
cXmnZK65rf*&DaaD
Yuc8$RikA34%ZoPPao98t
K8#mP2xR@vL9qN4zT6yU
```

### 3. Use character diversity
- ✓ Uppercase letters (A-Z)
- ✓ Lowercase letters (a-z)
- ✓ Numbers (0-9)
- ✓ Symbols (!@#$%^&*()_+-=[]{}|;:,.<>?)

## Installation

### Prerequisites
- Go 1.16 or higher

### Build from Source

```bash
git clone https://github.com/andinianst94/go-converter.git
cd go-converter
go build -o devtools
```

### Run the Program

```bash
./devtools        # Linux/macOS
devtools.exe  # Windows
```

## Usage

When you run the program, you'll see the main menu:

```
===Welcome to DevTools===
Choose an option:
1. Password Generator
2. Base64 Converter
Enter your choice (1 or 2):
```

### Option 1: Password Generator

After selecting option 1, you'll see:

```
===Password Generator===
Choose an option:
1. Generate Password
2. Generate Password with Custom Length
Enter your choice (1 or 2):
```

#### 1.1 Generate Password (Default)
Generates **10 strong random passwords** with 20 characters each.

**Example Output:**
```
✓ Generated 10 Random Passwords (length: 20 each):
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
   1. K8#mP2xR@vL9qN4zT6yU
   2. F5$aB7hC!wD3sE1jG9pM
   3. X2&nY4kZ@rT6mL8vP0cQ
   4. W9*eR3tU#oI5yA7sD2fG
   5. H4^jK6lM$xC8vB1nN3mZ
   6. Q7&wE9rT!yU2iP4oA6sD
   7. V3*fG5hJ#kL7zX9cV1bN
   8. M8@nB2vC$xZ4aS6dF0gH
   9. J5!qW7eR&tY9uI3oP1aS
  10. D2#fG4hJ*kL6mN8bV0cX
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

✓ These passwords are STRONG because they:
  1. Are at least 16 characters long
  2. Contain mixed-case letters, numbers and symbols
  3. Are completely random
```

#### 1.2 Generate Password with Custom Length
Allows you to specify the password length (minimum 16 characters) and generates **11 passwords**.

**Example:**
```
Enter desired password length (minimum 16): 24
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
   1. Z7|KCyLq@9:${mwrk@l;
   2. X>Y$vP,9ts)Qyp6bB5+c
   3. $4Rj(v?l5m<Tq7@P68U!
   4. ipG0vo+QB(sKFe4cIfEY
   5. vWA7j@Fi(I]3X{BheD5-
   6. Ui,A>#85dD},WnR@N(JN
   7. 4@Wi?EnFB]#A#R@om^(h
   8. b^@+#Nq?iHCf]%l.ksq8
   9. D_PQz#DQ6eXSuP<8V96X
  10. NR32iHN}|@%Q7vMj4MFM
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
```

If you enter a length less than 16:
```
Enter desired password length (minimum 16): 10
✗ Password length must be at least 16 characters for security.
  Generating with minimum length of 16 instead...
```

### Option 2: Base64 Converter

After selecting option 2, you'll see:

```
===Base 64 Converter===
Choose an option:
1. Encode
2. Decode
Enter your choice (1 or 2):
```

#### 2.1 Encode Plain Text to Base64

**Example:**
```
Enter your choice (1 or 2): 1
Enter plain text to encode: Hello World

Encoded Base64:
SGVsbG8gV29ybGQ=
```

#### 2.2 Decode Base64 to Plain Text

**Example:**
```
Enter your choice (1 or 2): 2
Enter Base64 text to decode: SGVsbG8gV29ybGQ=

Decoded plain text:
Hello World
```

**Invalid Base64 Example:**
```
Enter Base64 text to decode: Invalid!!!Base64
Error decoding Base64: illegal base64 data at input byte 7
Make sure the input is valid Base64 text.
```

## Project Structure

```
go-converter/
├── main.go                 # Main application entry point
├── password/
│   ├── password.go        # Password generation functions
│   └── helpers.go         # Helper functions and constants
├── base64/
│   └── base64.go          # Base64 encoding/decoding functions
├── go.mod                 # Go module file
└── README.md              # This file
```

## Security Notes

- This tool uses Go's `crypto/rand` package for cryptographically secure random number generation
- All passwords are generated locally on your machine
- No passwords are stored, transmitted, or logged anywhere
- Always store your passwords in a secure password manager
- Never reuse passwords across different services
- Change your passwords regularly, especially for sensitive accounts

## Best Practices for Passwords

### ✓ DO:
- Use unique passwords for each account
- Use a password manager to store passwords securely
- Enable two-factor authentication (2FA) when available
- Change passwords if you suspect a breach

### ✗ DON'T:
- Don't use personal information (names, birthdays, etc.)
- Don't use common words or phrases
- Don't share passwords with others
- Don't write passwords on paper or in unsecured files
- Don't use the same password for multiple accounts

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License

## Author

[Andini Anissa](https://github.com/andinianst93)
