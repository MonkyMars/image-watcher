package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/radovskyb/watcher"
)

var (
	supportedFormats = []string{
		".jpg", ".jpeg", ".png",
	}
	base       = "/home/[user]/Pictures" // Replace with your actual path
	workQueue  = make(chan string, 10)   // Buffered queue
	numWorkers = 2                       // Max 2 conversions at a time
)

func main() {
	w := watcher.New()
	w.FilterOps(watcher.Create)

	// Start the workers
	for i := range numWorkers {
		go worker(i)
	}

	// Watcher event loop
	go func() {
		for event := range w.Event {
			if event.IsDir() {
				// Automatically add new directories to the watcher
				if err := w.AddRecursive(event.Path); err != nil {
					log.Println("Failed to watch new dir:", err)
				}
				continue
			}

			ext := strings.ToLower(filepath.Ext(event.Path))
			if isSupportedFormat(ext) {
				workQueue <- event.Path
			}
		}
	}()

	// Error handling
	go func() {
		for err := range w.Error {
			log.Println("Watcher error:", err)
		}
	}()

	// Initial directory load
	if err := w.AddRecursive(base); err != nil {
		log.Fatalln(err)
	}

	// Start watching
	if err := w.Start(3 * time.Second); err != nil {
		log.Fatalln(err)
	}
}

// worker handles image conversion
func worker(id int) {
	for path := range workQueue {
		convert(path, id)
	}
}

func convert(path string, workerID int) {
	time.Sleep(2 * time.Second) // Let the file settle

	webpPath := strings.TrimSuffix(path, filepath.Ext(path)) + ".webp"
	cmd := exec.Command("cwebp", "-mt", "-q", "80", path, "-o", webpPath)
	cmd.Stderr = os.Stderr
	cmd.Stdout = io.Discard

	start := time.Now()
	if err := cmd.Run(); err != nil {
		log.Printf("[Worker %d] Conversion error: %v\n", workerID, err)
		return
	}

	if err := os.Remove(path); err != nil {
		log.Printf("[Worker %d] Failed to delete original: %v\n", workerID, err)
		return
	}

	rel, _ := filepath.Rel(base, webpPath)
	log.Printf("[Worker %d] Converted to %s in %.1fs\n", workerID, rel, time.Since(start).Seconds())
}

func isSupportedFormat(ext string) bool {
	return slices.Contains(supportedFormats, strings.ToLower(ext))
}
