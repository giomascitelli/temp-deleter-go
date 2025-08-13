# temp-deleter-go v2.1

A fast and efficient temporary file cleaner written in Go. Safely removes temporary files from Windows, Linux, and macOS systems with concurrent processing and detailed logging.

## Why use temp-deleter?

1. **Performance**: Temporary files slow down your computer and waste disk space unnecessarily.
2. **Privacy**: Temporary files may contain sensitive cached data. Regular cleanup helps protect your privacy.

## VirusTotal Scan Results

`temp-deleter-windows-amd64.exe`: https://www.virustotal.com/gui/file/d4ecfa978bbcad40ce598d8fb13f2c97f99e3275c5ef7560b862c5f2ff7cf442/details

`temp-deleter-windows-arm64.exe`: https://www.virustotal.com/gui/file/6f8b69e7d11ffdf2e40097d1bc5e0e90ba77cad8ccf72640946193f3f2d6b97e/detection

`temp-deleter-linux-amd64`: https://www.virustotal.com/gui/file/9b0cd21840b41cec0fc380d01e55ef949660bd2ed2dfe5b6015df2e49fb90eaa?nocache=1

`temp-deleter-linux-arm64`: https://www.virustotal.com/gui/file/aba0376e59d937f653622f92535e99e027174dd9713ccc52d4437d2637de8e66/detection

`temp-deleter-darwin-amd64`: https://www.virustotal.com/gui/file/893539aa9795c5ca44a06fcd127062e557f35963a481408b7cc38cb248ccd7aa/detection

`temp-deleter-darwin-arm64`: https://www.virustotal.com/gui/file/5456404a036cb7159e1fd007117b4e328e5e059b155e0b19c0e9fc22a11b0879?nocache=1

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

| Feature | Python v1.2 | Go v2.1 |
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
