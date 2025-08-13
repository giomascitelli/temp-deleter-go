// Package logger provides structured logging functionality with file output support.
// It wraps logrus with custom methods for temp-deleter operations.
package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// Logger wraps logrus with custom functionality
type Logger struct {
	*logrus.Logger
	file *os.File
}

// New creates a new logger instance with file output and structured formatting
func New(filename string) *Logger {
	log := logrus.New()

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		// Fallback to stdout if file creation fails
		log.SetOutput(os.Stdout)
		log.Errorf("Failed to create log file %s: %v", filename, err)
		return &Logger{Logger: log, file: nil}
	}

	log.SetOutput(file)
	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors:   true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	})

	log.WithFields(logrus.Fields{
		"version": "2.0.0",
		"time":    time.Now().Format(time.RFC3339),
	}).Info("Temp Deleter started")

	return &Logger{Logger: log, file: file}
}

// Close closes the log file
func (l *Logger) Close() {
	if l.file != nil {
		l.Info("Temp Deleter session ended")
		l.file.Close()
	}
}

// LogDeletion logs a successful file deletion
func (l *Logger) LogDeletion(path string, size int64, isDir bool) {
	fileType := "file"
	if isDir {
		fileType = "directory"
	}

	l.WithFields(logrus.Fields{
		"path": path,
		"size": size,
		"type": fileType,
	}).Info("Deleted")
}

// LogError logs an error during deletion
func (l *Logger) LogError(path string, err error) {
	l.WithFields(logrus.Fields{
		"path":  path,
		"error": err.Error(),
	}).Error("Failed to delete")
}

// LogSkipped logs a skipped item
func (l *Logger) LogSkipped(path string, reason string) {
	l.WithFields(logrus.Fields{
		"path":   path,
		"reason": reason,
	}).Warn("Skipped")
}

// LogDirectoryProcessing logs directory processing start
func (l *Logger) LogDirectoryProcessing(dir string) {
	l.WithField("directory", dir).Info("Processing directory")
}

// Infof provides formatted info logging
func (l *Logger) Infof(format string, args ...interface{}) {
	l.Info(fmt.Sprintf(format, args...))
}

// Errorf provides formatted error logging
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Error(fmt.Sprintf(format, args...))
}

// Warnf provides formatted warning logging
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Warn(fmt.Sprintf(format, args...))
}
