
# URL Shortener CLI

A powerful and user-friendly command-line tool built in Go to shorten URLs using various services. It features a sleek interface with colored output, history tracking, QR code generation, and support for multiple shortening services.

## Features
- Shorten URLs using services like TinyURL (with support for more services planned).
- Track and display a history of shortened URLs.
- Generate QR codes for shortened URLs.
- Command-line flags for easy configuration and usage.
- Stylish CLI with ASCII art, colors, and a progress spinner for a better user experience.

## Prerequisites
- [Go](https://golang.org/dl/) (version 1.16 or later recommended).
- Git (for cloning the repository).

The following external libraries are used and will be installed automatically during setup:
- `github.com/fatih/color` for colored output.
- `github.com/briandowns/spinner` for the progress spinner.
- `github.com/skip2/go-qrcode` for QR code generation.
- `github.com/common-nighthawk/go-figure` for ASCII art.

**Note for Linux users**: You may need `xclip` or `xsel` for clipboard functionality if added in future updates.

## Installation
Follow these steps to set up the project on your machine:

1. **Clone the repository**:
   ```bash
   git clone https://github.com/MRQ67/url-shorten.git
   cd url-shortener-cli
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **(Optional) Build the executable**:
   ```bash
   go build -o urlshort
   ```

## Usage
Run the tool using `go run main.go` or the built executable (`./urlshort`). It supports both interactive input and command-line flags.

### Command-Line Flags
- `-url <URL>`: Specify the URL to shorten.
- `-history`: Display the history of shortened URLs.
- `-qr`: Generate a QR code for the shortened URL.
- `-help`: Show usage information.

### Examples
- **Shorten a URL interactively**:
  ```bash
  go run main.go
  # Enter URL to shorten: https://example.com
  # Output: Shortened URL: https://tinyurl.com/abc
  ```

- **Shorten a URL with flags and generate a QR code**:
  ```bash
  go run main.go -url https://example.com -qr
  # Output: Shortened URL: https://tinyurl.com/abc
  #         QR code saved as qr.png
  ```

- **Display history**:
  ```bash
  go run main.go -history
  # Output: 1. https://example.com -> https://tinyurl.com/abc (2023-10-01T12:00:00Z)
  ```

- **Show help**:
  ```bash
  go run main.go -help
  # Output: Usage: urlshort [options]
  #         Options: [list of flags]
  ```
  
### Examples Output
```
   __  __    ____     __      _____    __                    __
  / / / /   / __ \   / /     / ___/   / /_   ____    _____  / /_  ___    ____   ___    _____
 / / / /   / /_/ /  / /      \__ \   / __ \ / __ \  / ___/ / __/ / _ \  / __ \ / _ \  / ___/
/ /_/ /   / _, _/  / /___   ___/ /  / / / // /_/ / / /    / /_  /  __/ / / / //  __/ / /
\____/   /_/ |_|  /_____/  /____/  /_/ /_/ \____/ /_/     \__/  \___/ /_/ /_/ \___/ /_/
-----------------------------------------
Enter URL to shorten: https://www.youtube.com/
Shortening...
Shortened URL: https://tinyurl.com/6gvvov (copied to clipboard)
Shortened URL: https://tinyurl.com/6gvvov
```

## Configuration
- **History File**: Shortened URLs are stored in `history.json` in the current directory. You can change the file path by editing the code if desired.

## Contributing
Contributions are welcome! To contribute:
1. Fork the repository.
2. Create a new branch:
   ```bash
   git checkout -b feature/your-feature
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add your feature"
   ```
4. Push to your branch:
   ```bash
   git push origin feature/your-feature
   ```
5. Open a Pull Request on GitHub.

Please ensure your code aligns with the project's style and include tests where applicable.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments
- Built with [Go](https://golang.org/).
- URL shortening powered by [TinyURL](https://tinyurl.com/).
- Libraries: [color](https://github.com/fatih/color), [spinner](https://github.com/briandowns/spinner), [go-qrcode](https://github.com/skip2/go-qrcode), [go-figure](https://github.com/common-nighthawk/go-figure).


