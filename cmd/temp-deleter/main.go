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

	// Initialize configuration
	cfg := config.New()

	// Initialize logger
	log := logger.New("temp_deleter.log")
	defer log.Close()

	// Initialize cleaner (false = not dry run)
	clean := cleaner.New(log, false)

	// Initialize storage (for potential cloud logging)
	store := storage.New("", log) // Empty SAS URL means no cloud storage

	// Get directories to clean
	directories := cfg.GetTempDirectories()

	log.Infof("Found %d directories to clean", len(directories))
	for _, dir := range directories {
		log.Infof("Target directory: %s", dir)
	}

	// Ask for confirmation
	fmt.Printf("\nFound %d temporary directories to clean.\n", len(directories))
	fmt.Print("Do you want to proceed? (y/N): ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	response := scanner.Text()

	if response != "y" && response != "Y" && response != "yes" && response != "Yes" {
		fmt.Println("Operation cancelled.")
		return
	}

	// Perform cleanup
	fmt.Println("\n🧹 Starting cleanup...")
	results := clean.CleanDirectories(directories)

	// Display results
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

	// Try to upload log to cloud storage if enabled
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
