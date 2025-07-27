# Image Watcher

An efficient Go-based utility that automatically monitors directories for new images and converts them to the modern .webp format, helping you optimize disk usage and improve web performance.

## Features

- **Automatic Directory Monitoring**: Polls target directories every 2 seconds for new image files
- **Concurrent Processing**: Utilizes Go routines for fast, efficient image conversion
- **Multiple Format Support**: Converts JPEG, PNG, and other common image formats to WebP
- **Space Optimization**: Significantly reduces file sizes while maintaining visual quality
- **Lightweight**: Minimal resource usage with efficient polling mechanism
- **Cross-Platform**: Works on Linux and macOS

## Requirements

### System Requirements

- **Go**: Version 1.19 or higher
- **Operating System**: Linux (primary), macOS, or Windows
- **Memory**: Minimum 512MB RAM (more recommended for large batches)
- **Disk Space**: Sufficient space for both original and converted images during processing

### Dependencies

The application uses the following Go package:

```go
// Core dependency (automatically handled by go mod)
github.com/radovskyb/watcher      // File system watching (if used)
```

### System Dependencies

#### Linux (Fedora/RHEL)
```bash
# Install libwebp-tools
sudo dnf install libwebp-tools
# Install exiftran
sudo dnf install exiftran
```

#### Linux (Ubuntu/Debian)
```bash
# Install webp
sudo apt-get install webp
# Install exiftran
sudo apt-get install exiftran
```

#### macOS
```bash
# Install libwebp via Homebrew
brew install webp
# Install exiftran via Homebrew
brew install exiftran
```

#### Windows
The current implementation is not tested on Windows, since I dont have a windows machine.
If you have tested it on Windows 10+, please open an issue and this will be updated.

## Installation

```bash
# Clone the repository
git clone https://github.com/MonkyMars/image-watcher.git
cd image-watcher

# Initialize Go module and install dependencies
go mod tidy

### make sure you've changed the path in main.go to your desired folder before building!

# Build the application
go build -o image-watcher

# Make executable (Linux/macOS)
chmod +x image-watcher
```
## Usage

```bash
# Monitor directory specified in main.go
./image-watcher
```
## Supported Image Formats

### Input Formats
- JPEG (.jpg, .jpeg)
- PNG (.png)

### Output Format
- WebP (.webp) - Modern, efficient image format

## Performance Considerations

### Optimal Settings

- **Quality**: 80-85 provides excellent balance between size and quality
- **Polling Interval**: 2s is efficient for most use cases; reduce for high-frequency scenarios
- **Concurrent Processing**: Limited by CPU cores and available memory

### Resource Usage

- **CPU**: Moderate usage during conversion, minimal during monitoring
- **Memory**: ~10-200MB base usage, increases with concurrent conversions
- **Disk I/O**: Temporary spike during conversion process

## Development

### Building from Source

```bash
# Clone and navigate
git clone https://github.com/MonkyMars/image-watcher.git
cd image-watcher

# Install dependencies
go mod download

# Run tests
go test ./...

# Build for current platform
go build -o image-watcher

# Cross-compile for Linux
GOOS=linux GOARCH=amd64 go build -o image-watcher-linux
```

### Project Structure

```
image-watcher/
├── main.go              # Main script
├── go.mod              # Go module definition
├── go.sum              # Dependency checksums
├── README.md           # This file
└── LICENSE             # Project license
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

- Follow Go best practices and `gofmt` standards
- Add tests for new functionality
- Update documentation for new features
- Ensure cross-platform compatibility

## Troubleshooting

### Common Issues

**"libwebp not found" error:**
```bash
# Linux: Install development headers
sudo dnf install libwebp-devel  # Fedora
sudo apt install libwebp-dev    # Ubuntu
```

**High CPU usage:**
- Increase polling interval
- Limit the numworkers in main.go to 1

**Permission denied:**
- Ensure read/write permissions on target directories
- Run with appropriate user privileges

**Memory issues with large images:**
- Process images in smaller batches
- Increase system memory or reduce concurrent operations

## License

Apache License - see [LICENSE](https://github.com/MonkyMars/image-watcher/blob/main/LICENSE) file for details.

## Acknowledgments

- [radovskyb/watcher](https://github.com/radovskyb/watcher) - File system watching
- Go community for excellent image processing libraries

---

**Built with ❤️ by [MonkyMars](https://github.com/MonkyMars)**
