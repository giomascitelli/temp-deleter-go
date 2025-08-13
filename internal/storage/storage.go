// Package storage provides cloud storage abstraction for log file uploads.
// This simplified version provides interface compatibility without Azure dependencies.
package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Storage handles log file operations
type Storage struct {
	enabled bool
	logger  Logger
}

// Logger interface for storage operations
type Logger interface {
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
}

// New creates a new storage instance with optional cloud storage support
func New(sasURL string, logger Logger) *Storage {
	if sasURL == "" {
		logger.Infof("No SAS URL provided, cloud logging disabled")
		return &Storage{enabled: false, logger: logger}
	}

	logger.Infof("Cloud logging would be enabled (Azure integration available in full version)")
	return &Storage{
		enabled: false, // Simplified version without Azure dependencies
		logger:  logger,
	}
}

// IsEnabled returns whether cloud storage is enabled
func (s *Storage) IsEnabled() bool {
	return s.enabled
}

// UploadLogFile uploads a log file to cloud storage (placeholder implementation)
func (s *Storage) UploadLogFile(logFilePath string) error {
	if !s.enabled {
		s.logger.Warnf("Cloud storage not enabled, log file saved locally: %s", logFilePath)
		return nil
	}

	return fmt.Errorf("cloud storage not implemented in this version")
}

// GenerateBlobName generates a unique timestamped name for log files
func (s *Storage) GenerateBlobName(baseName string) string {
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	hostname, _ := os.Hostname()
	if hostname == "" {
		hostname = "unknown"
	}

	baseName = filepath.Base(baseName)
	if baseName == "" {
		baseName = "temp-deleter"
	}

	return fmt.Sprintf("%s_%s_%s.log", hostname, baseName, timestamp)
}

// TestConnection tests cloud storage connectivity (placeholder implementation)
func (s *Storage) TestConnection() error {
	if !s.enabled {
		s.logger.Infof("Cloud storage connection test skipped (not enabled)")
		return nil
	}

	return fmt.Errorf("cloud storage not implemented in this version")
}
