# Temp Deleter Go

A fast and efficient temporary file cleaner written in Go, designed to safely remove temporary files from Windows, Linux, and macOS systems.

## Features

- **Cross-platform support**: Works on Windows, Linux, and macOS
- **Concurrent processing**: Uses multiple workers for faster file deletion
- **Safe deletion**: Skips system files and files currently in use
- **Detailed logging**: Comprehensive logging with structured output
- **User confirmation**: Always asks for confirmation before deletion
- **Size reporting**: Shows total space freed in human-readable format
- **Error handling**: Graceful error handling with detailed error reporting
- **Cloud logging**: Optional Azure Blob Storage integration for log uploads

## Supported Directories

### Windows
- `C:\Windows\Temp`
- `%LOCALAPPDATA%\Temp`
- `%TEMP%`
- Recycle Bin folders
- Windows Update cache
- Browser caches

### Linux
- `/tmp`
- `/var/cache`
- `~/.cache`
- `~/.local/share/Trash`
- Package manager caches

### macOS
- `~/Library/Caches`
- `~/.Trash`
- `/tmp`
- `/var/folders` (system temp)
- Browser caches

## Download & Installation

### 📦 Pre-built Binaries

Download the correct binary for your platform from the [GitHub Releases](https://github.com/giomascitelli/temp-deleter-go/releases) page:

| Platform | Architecture | Download | Size |
|----------|-------------|----------|------|
| **Windows 64-bit** | amd64 | `temp-deleter-windows-amd64.exe` | ~2.0 MB |
| **Windows 32-bit** | 386 | `temp-deleter-windows-386.exe` | ~1.9 MB |
| **Linux 64-bit** | amd64 | `temp-deleter-linux-amd64` | ~1.8 MB |
| **Linux 32-bit** | 386 | `temp-deleter-linux-386` | ~1.7 MB |
| **Linux ARM64** | arm64 | `temp-deleter-linux-arm64` | ~1.8 MB |
| **macOS Intel** | amd64 | `temp-deleter-macos-amd64` | ~1.8 MB |
| **macOS Apple Silicon** | arm64 | `temp-deleter-macos-arm64` | ~1.8 MB |

### 🚀 Platform-Specific Instructions

#### **Windows Users:**
1. Download `temp-deleter-windows-amd64.exe` (most users) or `temp-deleter-windows-386.exe` (32-bit systems)
2. Double-click to run - no installation required!

#### **Linux Users:**
1. Download the appropriate binary (`temp-deleter-linux-amd64` for Intel/AMD processors)
2. Make it executable:
   ```bash
   chmod +x temp-deleter-linux-amd64
   ```
3. Run it:
   ```bash
   ./temp-deleter-linux-amd64
   ```

#### **macOS Users:**
1. Download `temp-deleter-macos-amd64` (Intel Mac) or `temp-deleter-macos-arm64` (M1/M2 Mac)
2. Make it executable:
   ```bash
   chmod +x temp-deleter-macos-amd64
   ```
3. Run it:
   ```bash
   ./temp-deleter-macos-amd64
   ```

### 🛠️ Build from Source
```bash
git clone https://github.com/giomascitelli/temp-deleter-go.git
cd temp-deleter-go
go build ./cmd/temp-deleter
```

## Usage

The application provides a simple, interactive experience:

1. **Launch**: Run the executable for your platform
2. **Review**: It will show you the temporary directories it found
3. **Confirm**: Type `y` to proceed or `n` to cancel
4. **Clean**: Watch as it safely removes temporary files
5. **Results**: View a detailed report of what was cleaned

**Example output:**
```
🧹 Temp Deleter v2.0.0 - Fast temporary file cleaner
===============================================

Found 5 temporary directories to clean.
Do you want to proceed? (y/N): y

🧹 Starting cleanup...

✅ Cleanup completed!
📊 Results:
   Files processed: 1,234 (deleted: 1,180, failed: 12, skipped: 42)
   Directories processed: 56 (deleted: 52, failed: 1, skipped: 3)
   Space freed: 2.3 GB
   Processing time: 3.2s

🎉 All done! Press Enter to exit...
```

## Safety Features

- **File-in-use detection**: Skips files that are currently being used
- **System file protection**: Avoids deleting critical system files
- **User confirmation**: Always requires explicit user consent
- **Dry-run capability**: Can be configured to show what would be deleted without actually deleting
- **Detailed logging**: Every action is logged for audit purposes

## Performance & Advantages

This Go version offers significant performance improvements over the original Python version:

- **Concurrent processing**: Multiple workers process directories in parallel
- **Efficient memory usage**: Go's garbage collector and efficient memory management
- **Fast startup**: Compiled binary starts instantly (~0.1s vs 2-3s for Python)
- **Cross-platform native code**: No runtime dependencies required
- **Single executable**: No need to install Python, packages, or interpreters
- **Small footprint**: ~2MB download vs 50MB+ for Python with dependencies
- **Platform-optimized**: Each binary is compiled specifically for its target platform

## Configuration

The application uses sensible defaults and requires no configuration for basic usage. Advanced users can modify the source code to:

- Add custom temporary directories
- Adjust the number of worker goroutines
- Enable Azure Blob Storage logging
- Customize file filtering rules

## Logging

All operations are logged to `temp_deleter.log` in the current directory. The log includes:

- Start and end times
- Directories processed
- Files and directories deleted
- Errors encountered
- Total statistics

## Building

### Requirements
- Go 1.21 or later

### Build for Current Platform
```bash
go build ./cmd/temp-deleter
```

### Cross-Platform Building

Build for all platforms at once:

**On Windows:**
```batch
# Create dist directory
mkdir dist

# Windows builds
set GOOS=windows
set GOARCH=amd64
go build -ldflags "-s -w" -o dist/temp-deleter-windows-amd64.exe ./cmd/temp-deleter

set GOARCH=386
go build -ldflags "-s -w" -o dist/temp-deleter-windows-386.exe ./cmd/temp-deleter

# Linux builds
set GOOS=linux
set GOARCH=amd64
go build -ldflags "-s -w" -o dist/temp-deleter-linux-amd64 ./cmd/temp-deleter

set GOARCH=arm64
go build -ldflags "-s -w" -o dist/temp-deleter-linux-arm64 ./cmd/temp-deleter

# macOS builds
set GOOS=darwin
set GOARCH=amd64
go build -ldflags "-s -w" -o dist/temp-deleter-macos-amd64 ./cmd/temp-deleter

set GOARCH=arm64
go build -ldflags "-s -w" -o dist/temp-deleter-macos-arm64 ./cmd/temp-deleter
```

**On Linux/macOS:**
```bash
# Create dist directory
mkdir -p dist

# Build for all platforms
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o dist/temp-deleter-windows-amd64.exe ./cmd/temp-deleter
GOOS=windows GOARCH=386 go build -ldflags "-s -w" -o dist/temp-deleter-windows-386.exe ./cmd/temp-deleter
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o dist/temp-deleter-linux-amd64 ./cmd/temp-deleter
GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o dist/temp-deleter-linux-arm64 ./cmd/temp-deleter
GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o dist/temp-deleter-macos-amd64 ./cmd/temp-deleter
GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o dist/temp-deleter-macos-arm64 ./cmd/temp-deleter
```

### Build Flags Explained
- `-ldflags "-s -w"`: Strip debug information and symbol table for smaller binaries
- `-o`: Specify output filename
- `GOOS`: Target operating system (windows, linux, darwin)
- `GOARCH`: Target architecture (amd64, 386, arm64)

## Security

This tool only deletes files from well-known temporary directories. It will never:
- Delete files from user documents
- Delete system files
- Delete files from program directories
- Delete files without user confirmation

## Version History

- **v2.0**: Complete rewrite in Go for better performance
- **v1.2**: Original Python version with Azure integration
- **v1.1**: Added Linux support
- **v1.0**: Initial Windows-only version

## Comparison with Python Version

| Feature | Python v1.2 | Go v2.0 |
|---------|-------------|---------|
| Startup Time | ~2-3 seconds | ~0.1 seconds |
| Memory Usage | ~20-50 MB | ~5-10 MB |
| Concurrency | Single-threaded | Multi-threaded |
| Dependencies | Python + packages | None (static binary) |
| File Processing | Sequential | Concurrent |
| Error Handling | Basic | Comprehensive |
