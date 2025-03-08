package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/mdp/qrterminal"
)

// Define command-line flag variables
var (
	url         string
	showHistory bool
	generateQR  bool
	help        bool
)

// HistoryEntry represents a single entry in the history file
type HistoryEntry struct {
	OriginalURL  string `json:"original_url"`
	ShortenedURL string `json:"shortened_url"`
	Timestamp    string `json:"timestamp"`
}

// init sets up custom usage message for the help flag
func init() {
	flag.StringVar(&url, "url", "", "URL to shorten")
	flag.BoolVar(&showHistory, "history", false, "Show history of shortened URLs")
	flag.BoolVar(&generateQR, "qr", false, "Generate QR code for the shortened URL")
	flag.BoolVar(&help, "help", false, "Show this help message")

	flag.Usage = func() {
		fmt.Println("Usage: urlshort [options]")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}
}

// main function
func main() {
	title := figure.NewFigure("URL Shortener", "big", true).String()
	color.Cyan(title)
	fmt.Println("-----------------------------------------")

	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	// Handle the history flag
	if showHistory {
		displayHistory()
		os.Exit(0)
	}

	if url == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter URL to shorten: ")
		var err error
		url, err = reader.ReadString('\n')
		if err != nil {
			color.Red("Error reading input: %v", err)
			os.Exit(1)
		}
		url = strings.TrimSpace(url)
	}

	// Ensure a URL is provided
	if url == "" {
		color.Red("Error: URL is required")
		flag.Usage()
		os.Exit(1)
	}

	color.Yellow("Shortening...")
	shortURL, err := shortenURL(url)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
		return
	}
	err = clipboard.WriteAll(shortURL)
	if err == nil {
		fmt.Println("Shortened URL:", shortURL, "(copied to clipboard)")
	} else {
		fmt.Println("Shortened URL:", shortURL)
		fmt.Println("Note: Could not copy to clipboard:", err)
	}
	// Save to history
	saveToHistory(url, shortURL)

	// Display the shortened URL
	color.Green("Shortened URL: %s", shortURL)

	// Generate QR code if requested
	if generateQR {
		color.Yellow("Generating QR code...")
		generateQRCode(shortURL)
	}
}

// shortenURL shortens the given URL
func shortenURL(longURL string) (string, error) {
	url := fmt.Sprintf("http://tinyurl.com/api-create.php?url=%s", longURL)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// saveToHistory saves the shortened URL to history
func saveToHistory(originalURL, shortenedURL string) {
	history, err := readHistory()
	if err != nil {
		fmt.Println("Warning: Could not read history:", err)
		return
	}
	entry := HistoryEntry{
		OriginalURL:  originalURL,
		ShortenedURL: shortenedURL,
		Timestamp:    time.Now().UTC().Format(time.RFC3339),
	}
	history = append(history, entry)
	err = writeHistory(history)
	if err != nil {
		fmt.Println("Warning: Could not save to history:", err)
	}
}

// readHistory reads the history from history.json
func readHistory() ([]HistoryEntry, error) {
	data, err := os.ReadFile("history.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []HistoryEntry{}, nil
		}
		return nil, err
	}
	var history []HistoryEntry
	err = json.Unmarshal(data, &history)
	return history, err
}

// writeHistory writes the history to history.json
func writeHistory(history []HistoryEntry) error {
	data, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("history.json", data, 0644)
}

// displayHistory shows the list of previously shortened URLs
func displayHistory() {
	history, err := readHistory()
	if err != nil {
		color.Red("Error reading history: %v", err)
		return
	}
	for i, entry := range history {
		fmt.Printf("%d. %s -> %s (%s)\n", i+1, entry.OriginalURL, entry.ShortenedURL, entry.Timestamp)
	}
}

// generateQRCode creates a QR code for the given URL and saves it as qr.png
func generateQRCode(url string) {
	config := qrterminal.Config{
		Level:     qrterminal.L,
		Writer:    os.Stdout,
		BlackChar: qrterminal.BLACK,
		WhiteChar: qrterminal.WHITE,
		QuietZone: 1, // Smallest practical border
	}
	qrterminal.GenerateWithConfig(url, config)
	color.Green("QR code displayed above")
}
