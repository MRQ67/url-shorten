package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/atotto/clipboard"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a long URL: ")
	longURL, _ := reader.ReadString('\n')
	longURL = strings.TrimSpace(longURL)

	shortURL, err := shortenURL(longURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	err = clipboard.WriteAll(shortURL)
	if err == nil {
		fmt.Println("Shortened URL:", shortURL, "(copied to clipboard)")
	} else {
		fmt.Println("Shortened URL:", shortURL)
		fmt.Println("Note: Could not copy to clipboard:", err)
	}

}
