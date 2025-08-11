# Temp Deleter Go - Makefile

.PHONY: build clean test run install

# Build for current platform
build:
	go build -o temp-deleter ./cmd/temp-deleter

# Build for all platforms
build-all:
	GOOS=windows GOARCH=amd64 go build -o dist/temp-deleter-windows-amd64.exe ./cmd/temp-deleter
	GOOS=linux GOARCH=amd64 go build -o dist/temp-deleter-linux-amd64 ./cmd/temp-deleter
	GOOS=darwin GOARCH=amd64 go build -o dist/temp-deleter-darwin-amd64 ./cmd/temp-deleter
	GOOS=windows GOARCH=386 go build -o dist/temp-deleter-windows-386.exe ./cmd/temp-deleter
	GOOS=linux GOARCH=386 go build -o dist/temp-deleter-linux-386 ./cmd/temp-deleter

# Clean build artifacts
clean:
	rm -rf temp-deleter temp-deleter.exe dist/ temp_deleter.log

# Run tests
test:
	go test ./...

# Run the application
run:
	go run ./cmd/temp-deleter

# Install dependencies
deps:
	go mod download
	go mod tidy

# Create release directory
dist:
	mkdir -p dist

# Build with version info
build-release: dist
	go build -ldflags "-X main.Version=2.0.0 -s -w" -o dist/temp-deleter ./cmd/temp-deleter

# Show help
help:
	@echo "Available targets:"
	@echo "  build        - Build for current platform"
	@echo "  build-all    - Build for all platforms"
	@echo "  build-release - Build optimized release version"
	@echo "  clean        - Clean build artifacts"
	@echo "  test         - Run tests"
	@echo "  run          - Run the application"
	@echo "  deps         - Install dependencies"
	@echo "  help         - Show this help"
