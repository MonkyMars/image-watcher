# image-watcher

image-watcher is a Go-based utility that automatically monitors a directory for new images and converts them to the efficient .webp format to save storage space.

The application polls the target directory every 2 seconds, detecting newly added image files. When new images are found, it launches concurrent Go routines to convert each image to .webp, using Go routines for efficient and fast processing.

image-watcher helps optimize disk usage by converting images to a modern, web-optimized format automatically.
