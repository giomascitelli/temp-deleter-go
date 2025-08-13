# temp-deleter-go v2.0

A fast and efficient temporary file cleaner written in Go. Safely removes temporary files from Windows, Linux, and macOS systems with concurrent processing and detailed logging.

## Why use temp-deleter?

1. **Performance**: Temporary files slow down your computer and waste disk space unnecessarily.
2. **Privacy**: Temporary files may contain sensitive cached data. Regular cleanup helps protect your privacy.

## Supported Directories

**Windows**: `C:\Windows\Temp`, `%LOCALAPPDATA%\Temp`, Windows Update cache, Explorer cache, and more.

**Linux**: `/tmp`, `/var/cache`, `~/.cache`, package manager caches, and more.

**macOS**: `~/Library/Caches`, `~/.Trash`, `/tmp`, and more.

## Download & Usage

Download the correct binary for your platform from [GitHub Releases](https://github.com/giomascitelli/temp-deleter-go/releases):

- **Windows**: `temp-deleter-windows-amd64.exe` - Just double-click to run
- **Linux**: `temp-deleter-linux-amd64` - Make executable with `chmod +x` then run
- **macOS**: `temp-deleter-macos-amd64` (Intel) or `temp-deleter-macos-arm64` (Apple Silicon)

The application will show you what it found, ask for confirmation, then safely clean your temporary files.

## Features

- **Concurrent processing** with multiple workers for speed
- **Safe deletion** - skips system files and files in use
- **User confirmation** - always asks before deleting
- **Detailed logging** - comprehensive audit trail
- **Cross-platform** - Windows, Linux, and macOS support
- **No dependencies** - single executable, no runtime required

## Performance vs Python Version

| Feature | Python v1.2 | Go v2.0 |
|---------|-------------|---------|
| Startup Time | ~2-3 seconds | ~0.1 seconds |
| Memory Usage | ~20-50 MB | ~5-10 MB |
| Dependencies | Python + packages | None (static binary) |
| Processing | Sequential | Concurrent |

## Build from Source

```bash
git clone https://github.com/giomascitelli/temp-deleter-go.git
cd temp-deleter-go
go build ./cmd/temp-deleter
```

## For Advanced Users (Optional)

To enable Azure cloud logging for storing log files in the cloud, set the `AZURE_SAS_URL` environment variable with your container's SAS URL before running the application.
