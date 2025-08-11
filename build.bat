@echo off
echo Building Temp Deleter Go...
go build -o temp-deleter.exe ./cmd/temp-deleter
if %errorlevel% equ 0 (
    echo Build successful! You can now run temp-deleter.exe
) else (
    echo Build failed!
    pause
    exit /b 1
)
pause
