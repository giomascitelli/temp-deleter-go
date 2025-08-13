// Package config provides cross-platform configuration management for temp-deleter.
// It handles OS detection and provides platform-specific temporary directory lists.
package config

import (
	"os"
	"path/filepath"
	"runtime"
)

// Config holds application configuration
type Config struct {
	OS          string
	LogFile     string
	AzureSASURL string
}

// New creates a new configuration instance with platform-specific defaults
func New() *Config {
	return &Config{
		OS:          runtime.GOOS,
		LogFile:     "temp-deleter.log",
		AzureSASURL: os.Getenv("AZURE_SAS_URL"), // Optional: set via environment variable
	}
}

// GetTempDirectories returns the list of temporary directories for the current OS
func (c *Config) GetTempDirectories() []string {
	switch c.OS {
	case "windows":
		return c.getWindowsTempDirs()
	case "linux":
		return c.getLinuxTempDirs()
	case "darwin":
		return c.getMacOSTempDirs()
	default:
		return []string{}
	}
}

func (c *Config) getWindowsTempDirs() []string {
	userProfile := os.Getenv("USERPROFILE")
	if userProfile == "" {
		userProfile = os.Getenv("HOMEPATH")
	}

	dirs := []string{
		`C:\Windows\Temp`,
		`C:\Windows\Panther`,
		`C:\Windows\SoftwareDistribution\Download`,
	}

	if userProfile != "" {
		dirs = append(dirs,
			filepath.Join(userProfile, `AppData\Local\Temp`),
			filepath.Join(userProfile, `AppData\Local\Microsoft\Windows\Explorer`),
		)
	}

	return dirs
}

func (c *Config) getLinuxTempDirs() []string {
	homeDir := os.Getenv("HOME")

	dirs := []string{
		"/tmp",
		"/var/tmp",
		"/var/cache",
		"/var/cache/apt/archives",
		"/var/log",
	}

	if homeDir != "" {
		dirs = append(dirs,
			filepath.Join(homeDir, ".cache"),
			filepath.Join(homeDir, ".config"),
		)
	}

	return dirs
}

func (c *Config) getMacOSTempDirs() []string {
	homeDir := os.Getenv("HOME")

	dirs := []string{
		"/tmp",
		"/var/tmp",
		"/var/log",
	}

	if homeDir != "" {
		dirs = append(dirs,
			filepath.Join(homeDir, "Library/Caches"),
			filepath.Join(homeDir, "Library/Logs"),
			filepath.Join(homeDir, ".Trash"),
		)
	}

	return dirs
}
