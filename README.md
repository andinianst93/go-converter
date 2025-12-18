# Base64 Converter

A simple command-line tool written in Go for encoding plain text to Base64 and decoding Base64 to plain text.

## Features

- Interactive menu system
- Encode plain text to Base64
- Decode Base64 to plain text
- Error handling for invalid inputs
- Cross-platform support (Windows, Linux, macOS)

## Installation

### Download Pre-built Binary

Download the latest binary for your platform from the [Releases](https://github.com/yourusername/base64-converter/releases) page:

- **Windows**: `adcov-windows-amd64.exe`
- **Linux (amd64)**: `adcov-linux-amd64`
- **Linux (arm64)**: `adcov-linux-arm64`
- **macOS (Intel)**: `adcov-darwin-amd64`
- **macOS (Apple Silicon)**: `adcov-darwin-arm64`

### Build from Source

**Using Make (recommended):**
```bash
git clone https://github.com/yourusername/base64-converter.git
cd base64-converter

# Build for current platform
make build

# Build for all platforms
make build-all

# Install to /usr/local/bin
make install
```

**Using Go directly:**
```bash
git clone https://github.com/yourusername/base64-converter.git
cd base64-converter
go build -o bin/adcov base64_converter.go
```

## Usage

Run the program:

```bash
./adcov  # Linux/macOS
adcov.exe  # Windows
```

Follow the interactive prompts:

1. Choose option `1` to encode plain text to Base64
2. Choose option `2` to decode Base64 to plain text
3. Enter your text and press Enter

### Example

```
=== Base64 Converter ===
Choose an option:
1. Encode plain text to Base64
2. Decode Base64 to plain text
Enter your choice (1 or 2): 1

Enter plain text to encode: Hello World

Encoded Base64:
SGVsbG8gV29ybGQ=
```

## License

MIT License
