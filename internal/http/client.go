package http

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/JoaoPedr0Maciel/charm/internal/utils"
	"github.com/fatih/color"
	"github.com/tidwall/pretty"
)

// doRequest é uma função auxiliar interna que faz a requisição HTTP
func doRequest(method, url string, bearer *string, basic *string, contentType *string, data *string) (*http.Response, error) {
	startTime := time.Now()

	var bodyReader io.Reader
	if data != nil && *data != "" {
		bodyReader = strings.NewReader(*data)
	}

	req, err := http.NewRequest(method, url, bodyReader)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if bearer != nil && *bearer != "" {
		req.Header.Set("Authorization", "Bearer "+*bearer)
	}

	if basic != nil && *basic != "" {
		auth := base64.StdEncoding.EncodeToString([]byte(*basic))
		req.Header.Set("Authorization", "Basic "+auth)
	}

	// Define Content-Type (default: application/json se tiver data)
	if contentType != nil && *contentType != "" {
		req.Header.Set("Content-Type", *contentType)
	} else if data != nil && *data != "" {
		req.Header.Set("Content-Type", "application/json")
	}
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
	statusEmoji := utils.GetEmojiByStatusCode(statusCode)

	// Header principal
	fmt.Println()
	color.New(color.Bold, color.FgHiWhite).Println("╔═══════════════════════════════════════════════════════════════════════════")
	statusLine := fmt.Sprintf("║ %s  Status: %d  │  %s   │  time: %s", statusEmoji, statusCode, url, time.Since(startTime).Round(time.Millisecond))
	color.New(color.Bold, utils.GetColorByStatus(statusCode)).Println(statusLine)
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
	for _, line := range utils.SplitLines(bodyLines) {
		if line != "" {
			fmt.Println("║ " + line)
		}
	}

	color.New(color.Bold, color.FgHiWhite).Println("╚═══════════════════════════════════════════════════════════════════════════")
	fmt.Println()

	return resp, nil
}

func Get(url string, bearer *string, basic *string, contentType *string) (*http.Response, error) {
	return doRequest("GET", url, bearer, basic, contentType, nil)
}

func Post(url string, bearer *string, basic *string, contentType *string, data *string) (*http.Response, error) {
	return doRequest("POST", url, bearer, basic, contentType, data)
}

func Put(url string, bearer *string, basic *string, contentType *string, data *string) (*http.Response, error) {
	return doRequest("PUT", url, bearer, basic, contentType, data)
}

func Patch(url string, bearer *string, basic *string, contentType *string, data *string) (*http.Response, error) {
	return doRequest("PATCH", url, bearer, basic, contentType, data)
}

func Delete(url string, bearer *string, basic *string, contentType *string, data *string) (*http.Response, error) {
	return doRequest("DELETE", url, bearer, basic, contentType, data)
}
