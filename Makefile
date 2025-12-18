.PHONY: build build-all clean install help

BINARY_NAME=adcov
BUILD_DIR=bin
SOURCE_FILE=main.go

# Default build for current platform
build:
	@echo "Building $(BINARY_NAME) for current platform..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SOURCE_FILE)
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

# Build for all platforms
build-all:
	@echo "Building $(BINARY_NAME) for all platforms..."
	@mkdir -p $(BUILD_DIR)

	@echo "Building for Windows (amd64)..."
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(SOURCE_FILE)

	@echo "Building for Linux (amd64)..."
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(SOURCE_FILE)

	@echo "Building for Linux (arm64)..."
	GOOS=linux GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 $(SOURCE_FILE)

	@echo "Building for macOS (amd64)..."
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(SOURCE_FILE)

	@echo "Building for macOS (arm64)..."
	GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 $(SOURCE_FILE)

	@echo "All builds complete! Binaries in $(BUILD_DIR)/"
	@ls -lh $(BUILD_DIR)/

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete!"

# Install to /usr/local/bin (Unix-like systems only)
install: build
	@echo "Installing $(BINARY_NAME) to /usr/local/bin..."
	@sudo cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/
	@echo "Installation complete! You can now run '$(BINARY_NAME)' from anywhere."

# Run the program
run: build
	@./$(BUILD_DIR)/$(BINARY_NAME)

# Show help
help:
	@echo "Available targets:"
	@echo "  make build      - Build for current platform (default)"
	@echo "  make build-all  - Build for all platforms"
	@echo "  make clean      - Remove build artifacts"
	@echo "  make install    - Install to /usr/local/bin (requires sudo)"
	@echo "  make run        - Build and run the program"
	@echo "  make help       - Show this help message"
