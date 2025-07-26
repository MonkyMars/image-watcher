package main

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"

	"github.com/radovskyb/watcher"
)

var (
	supportedFormats []string = []string{
		".jpg",
		".png",
		".jpeg",
	}
	// Replace with your actual base directory
	base string = "/home/[user]/Pictures" // Base directory to watch
)

func main() {
	w := watcher.New()
	w.FilterOps(watcher.Create)

	go func() {
		for event := range w.Event {
			if event.IsDir() {
				continue
			}
			extension := strings.ToLower(filepath.Ext(event.Path))

			// Check if the file extension is supported
			if isSupportedFormat(extension) {
				go convert(event.Path) // Run conversion in its own goroutine
			}
		}
	}()

	go func() {
		for err := range w.Error {
			log.Println("Watcher error:", err)
		}
	}()

	if err := w.AddRecursive(base); err != nil {
		log.Fatalln(err)
	}

	if err := w.Start(time.Second * 2); err != nil {
		log.Fatalln(err)
	}
}

func convert(path string) {
	time.Sleep(time.Second) // Let the file settle
	start := time.Now()

	file, err := os.Open(path)
	if err != nil {
		log.Println("Open error:", err)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Println("Decode error:", err)
		return
	}

	webpPath := strings.TrimSuffix(path, filepath.Ext(path)) + ".webp"
	output, err := os.Create(webpPath)
	if err != nil {
		log.Println("Create output error:", err)
		return
	}
	defer output.Close()

	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 80)
	if err != nil {
		log.Println("WebP options error:", err)
		return
	}

	if err := webp.Encode(output, img, options); err != nil {
		log.Println("WebP encoding error:", err)
		return
	}

	if err := os.Remove(path); err != nil {
		log.Println("Failed to delete original file:", err)
	}

	relative, err := filepath.Rel(base, webpPath)
	if err != nil {
		log.Println("Failed to get relative path:", err)
	}

	elapsed := time.Since(start)
	seconds := elapsed.Seconds()
	log.Printf("Converted to %s in %.1f seconds\n", relative, seconds)
}

func isSupportedFormat(ext string) bool {
	return slices.Contains(supportedFormats, strings.ToLower(ext))
}
