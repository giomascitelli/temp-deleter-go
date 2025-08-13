// Package cleaner provides concurrent file and directory cleaning operations.
// It implements safe deletion with protection mechanisms and detailed result tracking.
package cleaner

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// Logger interface for cleaner operations
type Logger interface {
	LogDeletion(path string, size int64, isDir bool)
	LogError(path string, err error)
	LogSkipped(path string, reason string)
	LogDirectoryProcessing(dir string)
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
}

// CleanupResult holds the results of a cleanup operation
type CleanupResult struct {
	TotalFiles     int64
	TotalDirs      int64
	TotalSize      int64
	DeletedFiles   int64
	DeletedDirs    int64
	DeletedSize    int64
	FailedFiles    int64
	FailedDirs     int64
	SkippedFiles   int64
	SkippedDirs    int64
	ProcessingTime time.Duration
	ErrorMessages  []string
}

// Cleaner handles file and directory deletion operations
type Cleaner struct {
	logger     Logger
	maxWorkers int
	dryRun     bool
}

// New creates a new cleaner instance with optimal worker count based on CPU cores
func New(logger Logger, dryRun bool) *Cleaner {
	maxWorkers := runtime.NumCPU()
	if maxWorkers < 2 {
		maxWorkers = 2
	}
	if maxWorkers > 8 {
		maxWorkers = 8
	}

	return &Cleaner{
		logger:     logger,
		maxWorkers: maxWorkers,
		dryRun:     dryRun,
	}
}

// CleanDirectories cleans multiple directories concurrently using worker pool pattern
func (c *Cleaner) CleanDirectories(directories []string) *CleanupResult {
	startTime := time.Now()
	result := &CleanupResult{}

	c.logger.Infof("Starting cleanup of %d directories (workers: %d, dry-run: %v)",
		len(directories), c.maxWorkers, c.dryRun)

	workChan := make(chan string, len(directories))
	resultChan := make(chan *CleanupResult, len(directories))

	var wg sync.WaitGroup
	for i := 0; i < c.maxWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for dir := range workChan {
				dirResult := c.cleanDirectory(dir)
				resultChan <- dirResult
			}
		}()
	}

	for _, dir := range directories {
		workChan <- dir
	}
	close(workChan)

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for dirResult := range resultChan {
		c.mergeResults(result, dirResult)
	}

	result.ProcessingTime = time.Since(startTime)

	c.logger.Infof("Cleanup completed in %v", result.ProcessingTime)
	return result
}

func (c *Cleaner) cleanDirectory(dirPath string) *CleanupResult {
	result := &CleanupResult{}

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		c.logger.LogSkipped(dirPath, "directory does not exist")
		atomic.AddInt64(&result.SkippedDirs, 1)
		return result
	}

	c.logger.LogDirectoryProcessing(dirPath)

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			c.logger.LogError(path, err)
			if info != nil && info.IsDir() {
				atomic.AddInt64(&result.FailedDirs, 1)
			} else {
				atomic.AddInt64(&result.FailedFiles, 1)
			}
			result.ErrorMessages = append(result.ErrorMessages,
				fmt.Sprintf("Walk error for %s: %v", path, err))
			return filepath.SkipDir
		}

		// Skip the root directory itself
		if path == dirPath {
			return nil
		}

		if info.IsDir() {
			atomic.AddInt64(&result.TotalDirs, 1)
		} else {
			atomic.AddInt64(&result.TotalFiles, 1)
			atomic.AddInt64(&result.TotalSize, info.Size())
		}

		if c.shouldDelete(path, info) {
			if c.dryRun {
				c.logger.Infof("DRY RUN: Would delete %s", path)
				if info.IsDir() {
					atomic.AddInt64(&result.DeletedDirs, 1)
				} else {
					atomic.AddInt64(&result.DeletedFiles, 1)
					atomic.AddInt64(&result.DeletedSize, info.Size())
				}
			} else {
				if err := c.deleteItem(path, info); err != nil {
					c.logger.LogError(path, err)
					if info.IsDir() {
						atomic.AddInt64(&result.FailedDirs, 1)
					} else {
						atomic.AddInt64(&result.FailedFiles, 1)
					}
					result.ErrorMessages = append(result.ErrorMessages,
						fmt.Sprintf("Delete error for %s: %v", path, err))
				} else {
					c.logger.LogDeletion(path, info.Size(), info.IsDir())
					if info.IsDir() {
						atomic.AddInt64(&result.DeletedDirs, 1)
					} else {
						atomic.AddInt64(&result.DeletedFiles, 1)
						atomic.AddInt64(&result.DeletedSize, info.Size())
					}
				}
			}
		} else {
			c.logger.LogSkipped(path, "protected or system file")
			if info.IsDir() {
				atomic.AddInt64(&result.SkippedDirs, 1)
			} else {
				atomic.AddInt64(&result.SkippedFiles, 1)
			}
		}

		return nil
	})

	if err != nil {
		c.logger.LogError(dirPath, err)
		result.ErrorMessages = append(result.ErrorMessages,
			fmt.Sprintf("Directory walk error for %s: %v", dirPath, err))
	}

	return result
}

// shouldDelete determines if an item should be deleted based on safety rules
func (c *Cleaner) shouldDelete(path string, info os.FileInfo) bool {
	// Skip hidden files and directories on Unix-like systems
	if runtime.GOOS != "windows" && filepath.Base(path)[0] == '.' {
		return false
	}

	// Skip critical Windows system files
	if runtime.GOOS == "windows" {
		name := filepath.Base(path)
		if name == "desktop.ini" || name == "thumbs.db" || name == "Thumbs.db" {
			return false
		}
	}

	// Basic file-in-use check for files
	if !info.IsDir() {
		if file, err := os.OpenFile(path, os.O_WRONLY, 0); err == nil {
			// Handle close error (gosec G104)
			if closeErr := file.Close(); closeErr != nil {
				return false // Consider file in-use if we can't close it
			}
		} else {
			// If we can't open for writing, it might be in use
			return false
		}
	}

	return true
}

func (c *Cleaner) deleteItem(path string, info os.FileInfo) error {
	if info.IsDir() {
		return os.RemoveAll(path)
	}
	return os.Remove(path)
}

// mergeResults combines results from multiple worker goroutines
func (c *Cleaner) mergeResults(target, source *CleanupResult) {
	atomic.AddInt64(&target.TotalFiles, source.TotalFiles)
	atomic.AddInt64(&target.TotalDirs, source.TotalDirs)
	atomic.AddInt64(&target.TotalSize, source.TotalSize)
	atomic.AddInt64(&target.DeletedFiles, source.DeletedFiles)
	atomic.AddInt64(&target.DeletedDirs, source.DeletedDirs)
	atomic.AddInt64(&target.DeletedSize, source.DeletedSize)
	atomic.AddInt64(&target.FailedFiles, source.FailedFiles)
	atomic.AddInt64(&target.FailedDirs, source.FailedDirs)
	atomic.AddInt64(&target.SkippedFiles, source.SkippedFiles)
	atomic.AddInt64(&target.SkippedDirs, source.SkippedDirs)

	// Note: not thread-safe, but called after workers complete
	target.ErrorMessages = append(target.ErrorMessages, source.ErrorMessages...)
}

// FormatSize formats bytes into human-readable format
func FormatSize(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}
