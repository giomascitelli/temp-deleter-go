// Package main provides the entry point for the temp-deleter application.
// Temp Deleter is a cross-platform utility for safely cleaning temporary files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"temp-deleter/internal/cleaner"
	"temp-deleter/internal/config"
	"temp-deleter/internal/logger"
	"temp-deleter/internal/storage"
)

const Version = "2.0.0"

func main() {
	fmt.Printf("🧹 Temp Deleter v%s - Fast temporary file cleaner\n", Version)
	fmt.Println("===============================================")

	cfg := config.New()
	log := logger.New("temp_deleter.log")
	defer log.Close()

	// Disable dry-run mode for actual file deletion
	clean := cleaner.New(log, false)

	// Cloud storage is optional via SAS URL
	store := storage.New("", log)

	directories := cfg.GetTempDirectories()

	log.Infof("Found %d directories to clean", len(directories))
	for _, dir := range directories {
		log.Infof("Target directory: %s", dir)
	}

	fmt.Printf("\nFound %d temporary directories to clean.\n", len(directories))
	fmt.Print("Do you want to proceed? (y/N): ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	response := scanner.Text()

	if response != "y" && response != "Y" && response != "yes" && response != "Yes" {
		fmt.Println("Operation canceled.")
		return
	}

	fmt.Println("\n🧹 Starting cleanup...")
	results := clean.CleanDirectories(directories)

	// Display comprehensive results to user
	fmt.Printf("\n✅ Cleanup completed!\n")
	fmt.Printf("📊 Results:\n")
	fmt.Printf("   Files processed: %d (deleted: %d, failed: %d, skipped: %d)\n",
		results.TotalFiles, results.DeletedFiles, results.FailedFiles, results.SkippedFiles)
	fmt.Printf("   Directories processed: %d (deleted: %d, failed: %d, skipped: %d)\n",
		results.TotalDirs, results.DeletedDirs, results.FailedDirs, results.SkippedDirs)
	fmt.Printf("   Space freed: %s\n", cleaner.FormatSize(results.DeletedSize))
	fmt.Printf("   Processing time: %v\n", results.ProcessingTime)

	if len(results.ErrorMessages) > 0 {
		fmt.Printf("   ⚠️  %d errors occurred (check log file for details)\n", len(results.ErrorMessages))
	}

	// Upload logs to cloud storage if configured
	if store.IsEnabled() {
		fmt.Println("\n☁️  Uploading log to cloud storage...")
		if err := store.UploadLogFile("temp_deleter.log"); err != nil {
			log.Errorf("Failed to upload log file: %v", err)
			fmt.Printf("   ⚠️  Failed to upload log file: %v\n", err)
		} else {
			fmt.Println("   ✅ Log uploaded successfully")
		}
	}

	fmt.Println("\n🎉 All done! Press Enter to exit...")
	scanner.Scan()
}
