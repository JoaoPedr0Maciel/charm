package http

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/tidwall/pretty"
)

func Get(url string, authorization *string, contentType *string) (*http.Response, error) {
	startTime := time.Now()
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if authorization != nil {
		req.Header.Set("Authorization", *authorization)
	}

	if contentType != nil {
		req.Header.Set("Content-Type", *contentType)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	statusCode := resp.StatusCode
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error getting headers:", err)
		os.Exit(1)
	}

	formatted := pretty.Pretty(body)
	statusEmoji := getEmojiByStatusCode(statusCode)

	// Header principal
	fmt.Println()
	color.New(color.Bold, color.FgHiWhite).Println("╔═══════════════════════════════════════════════════════════════════════════")
	statusLine := fmt.Sprintf("║ %s  Status: %d  │  %s   │  time: %s", statusEmoji, statusCode, url, time.Since(startTime).Round(time.Millisecond))
	color.New(color.Bold, getColorByStatus(statusCode)).Println(statusLine)
	color.New(color.Bold, color.FgHiWhite).Println("╠═══════════════════════════════════════════════════════════════════════════")

	// Request Headers
	color.New(color.Bold, color.FgHiCyan).Println("║ 📋 HEADERS")
	color.New(color.FgHiWhite).Println("╟───────────────────────────────────────────────────────────────────────────")
	if contentType := req.Header.Get("Content-Type"); contentType != "" {
		color.New(color.FgYellow).Printf("║   • Content-Type: ")
		color.New(color.FgHiYellow).Println(contentType)
	}
	if auth := req.Header.Get("Authorization"); auth != "" {
		color.New(color.FgYellow).Printf("║   • Authorization: ")
		color.New(color.FgHiYellow).Println(auth)
	}

	// Body
	color.New(color.FgHiWhite).Println("╠═══════════════════════════════════════════════════════════════════════════")
	color.New(color.Bold, color.FgHiCyan).Println("║ 📦 RESPONSE BODY")
	color.New(color.FgHiWhite).Println("╟───────────────────────────────────────────────────────────────────────────")
	colored := pretty.Color(formatted, nil)

	// Adicionar prefixo em cada linha do body
	bodyLines := string(colored)
	for _, line := range splitLines(bodyLines) {
		if line != "" {
			fmt.Println("║ " + line)
		}
	}

	color.New(color.Bold, color.FgHiWhite).Println("╚═══════════════════════════════════════════════════════════════════════════")
	fmt.Println()

	return resp, nil
}

func getEmojiByStatusCode(statusCode int) string {
	if statusCode >= 200 && statusCode < 300 {
		return "✨"
	}

	if statusCode >= 300 && statusCode < 400 {
		return "↪️"
	}

	if statusCode >= 400 && statusCode < 500 {
		return "⚠️"
	}

	return "❌"
}

func getColorByStatus(statusCode int) color.Attribute {
	if statusCode >= 200 && statusCode < 300 {
		return color.FgHiGreen
	}

	if statusCode >= 300 && statusCode < 400 {
		return color.FgHiYellow
	}

	return color.FgHiRed

}

func splitLines(s string) []string {
	lines := []string{}
	currentLine := ""
	for _, char := range s {
		if char == '\n' {
			lines = append(lines, currentLine)
			currentLine = ""
		} else {
			currentLine += string(char)
		}
	}
	if currentLine != "" {
		lines = append(lines, currentLine)
	}
	return lines
}
