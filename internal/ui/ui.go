package ui

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/JoaoPedr0Maciel/charm/internal/structs"
	"github.com/fatih/color"
	"github.com/tidwall/pretty"
)

func Display(display structs.Display) {
	fmt.Println()
	DisplayHeader(display.Method, display.URL, display.Response, display.Body, display.TotalTime)
	DisplayRequest(display.Method, display.URL, display.Request, display.AuthHeader, display.Data)
	DisplayResponse(display.Response, display.Body, display.TotalTime)
	DisplayTiming(display.Timing, display.TotalTime)
}

func DisplayHeader(method, url string, resp *http.Response, body []byte, totalTime time.Duration) {
	statusEmoji := GetEmojiByStatusCode(resp.StatusCode)
	statusColor := GetColorByStatus(resp.StatusCode)

	cyan := color.New(color.FgHiCyan)
	white := color.New(color.FgHiWhite)
	gray := color.New(color.FgWhite)
	bold := color.New(color.Bold)

	white.Println("╭─────────────────────────────────────────────────────────────────────────────╮")
	cyan.Print("│ ✨ ")
	bold.Print("CHARM CLI")
	gray.Printf(" v1.0.0")
	fmt.Println(strings.Repeat(" ", 60) + "│")

	white.Println("├─────────────────────────────────────────────────────────────────────────────┤")

	// Método e URL
	cyan.Printf("│ 🚀 ")
	bold.Printf("%s", method)
	fmt.Print(" ")
	fmt.Println(truncateString(url, 65) + strings.Repeat(" ", max(0, 67-len(url))) + "│")

	// Status, Time, Size
	cyan.Printf("│ ")
	color.New(statusColor, color.Bold).Printf("%s %d %s", statusEmoji, resp.StatusCode, http.StatusText(resp.StatusCode))
	gray.Printf("  │  ⏱️  %s  │  📦 %s", totalTime.Round(time.Millisecond), FormatBytes(int64(len(body))))
	fmt.Println(strings.Repeat(" ", max(0, 30-len(totalTime.Round(time.Millisecond).String())-len(FormatBytes(int64(len(body)))))) + "│")

	white.Println("╰─────────────────────────────────────────────────────────────────────────────╯")
	fmt.Println()
}

func DisplayRequest(method, url string, req *http.Request, authHeader string, data string) {
	cyan := color.New(color.FgHiCyan)
	white := color.New(color.FgHiWhite)
	gray := color.New(color.FgWhite)
	yellow := color.New(color.FgYellow)

	cyan.Println("╭─ 📤 REQUEST ────────────────────────────────────────────────────────────────╮")

	yellow.Print("│ Method:   ")
	white.Println(method + strings.Repeat(" ", max(0, 65-len(method))) + "│")

	yellow.Print("│ URL:      ")
	white.Println(truncateString(url, 65) + strings.Repeat(" ", max(0, 65-len(url))) + "│")

	if req.Header.Get("Content-Type") != "" || authHeader != "" {
		yellow.Println("│ Headers:  " + strings.Repeat(" ", 65) + "│")

		if ct := req.Header.Get("Content-Type"); ct != "" {
			gray.Printf("│   • Content-Type: ")
			white.Println(ct + strings.Repeat(" ", max(0, 57-len(ct))) + "│")
		}

		if authHeader != "" {
			gray.Printf("│   • Authorization: ")
			maskedAuth := MaskToken(authHeader)
			white.Println(maskedAuth + strings.Repeat(" ", max(0, 56-len(maskedAuth))) + "│")
		}
	}

	if data != "" {
		yellow.Println("│ Body:     " + strings.Repeat(" ", 65) + "│")
		bodyPreview := truncateString(data, 65)
		gray.Printf("│   ")
		white.Println(bodyPreview + strings.Repeat(" ", max(0, 73-len(bodyPreview))) + "│")
	}

	white.Println("╰─────────────────────────────────────────────────────────────────────────────╯")
	fmt.Println()
}

func DisplayResponse(resp *http.Response, body []byte, totalTime time.Duration) {
	statusEmoji := GetEmojiByStatusCode(resp.StatusCode)
	statusColor := GetColorByStatus(resp.StatusCode)

	cyan := color.New(color.FgHiCyan)
	white := color.New(color.FgHiWhite)
	gray := color.New(color.FgWhite)
	yellow := color.New(color.FgYellow)
	green := color.New(color.FgHiGreen)

	cyan.Println("╭─ 📥 RESPONSE ───────────────────────────────────────────────────────────────╮")

	yellow.Print("│ Status:   ")
	color.New(statusColor).Printf("%s %d %s", statusEmoji, resp.StatusCode, http.StatusText(resp.StatusCode))
	fmt.Println(strings.Repeat(" ", max(0, 56-len(http.StatusText(resp.StatusCode)))) + "│")

	yellow.Print("│ Time:     ")
	green.Print(totalTime.Round(time.Millisecond))
	fmt.Println(strings.Repeat(" ", max(0, 65-len(totalTime.Round(time.Millisecond).String()))) + "│")

	yellow.Print("│ Size:     ")
	green.Print(FormatBytes(int64(len(body))))
	fmt.Println(strings.Repeat(" ", max(0, 65-len(FormatBytes(int64(len(body)))))) + "│")

	yellow.Println("│ Headers:  " + strings.Repeat(" ", 65) + "│")

	importantHeaders := []string{"Content-Type", "Content-Length", "Server", "Date", "X-Request-ID", "X-RateLimit-Remaining"}
	for _, headerName := range importantHeaders {
		if value := resp.Header.Get(headerName); value != "" {
			gray.Printf("│   • %s: ", headerName)
			headerValue := truncateString(value, 56-len(headerName))
			white.Println(headerValue + strings.Repeat(" ", max(0, 58-len(headerName)-len(headerValue))) + "│")
		}
	}

	white.Println("│" + strings.Repeat(" ", 79) + "│")

	yellow.Println("│ Body:     " + strings.Repeat(" ", 65) + "│")

	if len(body) > 0 {
		formatted := pretty.Pretty(body)
		colored := pretty.Color(formatted, nil)

		bodyLines := SplitLines(string(colored))
		for i, line := range bodyLines {
			if i > 20 {
				gray.Println("│   ... (truncated)" + strings.Repeat(" ", 58) + "│")
				break
			}
			if line != "" {
				truncated := truncateString(line, 75)
				fmt.Print("│ ")
				fmt.Print(truncated)
				fmt.Println(strings.Repeat(" ", max(0, 77-visualLen(line))) + "│")
			}
		}
	} else {
		gray.Println("│   (empty)" + strings.Repeat(" ", 66) + "│")
	}

	white.Println("╰─────────────────────────────────────────────────────────────────────────────╯")
	fmt.Println()
}

func DisplayTiming(timing *structs.TimingInfo, totalTime time.Duration) {
	magenta := color.New(color.FgHiMagenta)
	white := color.New(color.FgHiWhite)
	yellow := color.New(color.FgYellow)
	cyan := color.New(color.FgHiCyan)
	green := color.New(color.FgHiGreen)

	magenta.Println("╭─ 📊 TIMING ─────────────────────────────────────────────────────────────────╮")

	dnsTime := timing.DNSDone.Sub(timing.DNSStart)
	connectTime := timing.ConnectDone.Sub(timing.ConnectStart)
	tlsTime := timing.TLSDone.Sub(timing.TLSStart)
	transferTime := timing.ResponseDone.Sub(timing.ResponseStart)

	if !timing.DNSStart.IsZero() && !timing.DNSDone.IsZero() {
		yellow.Print("│   • DNS Lookup:    ")
		cyan.Print(dnsTime.Round(time.Millisecond))
		fmt.Println(strings.Repeat(" ", max(0, 54-len(dnsTime.Round(time.Millisecond).String()))) + "│")
	}

	if !timing.ConnectStart.IsZero() && !timing.ConnectDone.IsZero() {
		yellow.Print("│   • TCP Connect:   ")
		cyan.Print(connectTime.Round(time.Millisecond))
		fmt.Println(strings.Repeat(" ", max(0, 54-len(connectTime.Round(time.Millisecond).String()))) + "│")
	}

	if !timing.TLSStart.IsZero() && !timing.TLSDone.IsZero() {
		yellow.Print("│   • TLS Handshake: ")
		cyan.Print(tlsTime.Round(time.Millisecond))
		fmt.Println(strings.Repeat(" ", max(0, 54-len(tlsTime.Round(time.Millisecond).String()))) + "│")
	}

	if !timing.ResponseStart.IsZero() {
		yellow.Print("│   • Transfer:      ")
		cyan.Print(transferTime.Round(time.Millisecond))
		fmt.Println(strings.Repeat(" ", max(0, 54-len(transferTime.Round(time.Millisecond).String()))) + "│")
	}

	yellow.Print("│   • Total:         ")
	green.Print(totalTime.Round(time.Millisecond))
	fmt.Println(strings.Repeat(" ", max(0, 54-len(totalTime.Round(time.Millisecond).String()))) + "│")

	white.Println("╰─────────────────────────────────────────────────────────────────────────────╯")
	fmt.Println()
}

func FormatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

func MaskToken(token string) string {
	parts := strings.SplitN(token, " ", 2)
	if len(parts) != 2 {
		return token
	}

	tokenType := parts[0]
	tokenValue := parts[1]

	if len(tokenValue) <= 8 {
		return tokenType + " " + strings.Repeat("•", len(tokenValue))
	}

	return tokenType + " " + strings.Repeat("•", len(tokenValue)-4) + tokenValue[len(tokenValue)-4:]
}

func truncateString(s string, maxLen int) string {
	cleaned := stripAnsiCodes(s)
	if len(cleaned) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

func stripAnsiCodes(s string) string {
	result := s
	for strings.Contains(result, "\x1b[") {
		start := strings.Index(result, "\x1b[")
		end := strings.Index(result[start:], "m")
		if end == -1 {
			break
		}
		result = result[:start] + result[start+end+1:]
	}
	return result
}

func visualLen(s string) int {
	return len(stripAnsiCodes(s))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func GetEmojiByStatusCode(statusCode int) string {
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

func GetColorByStatus(statusCode int) color.Attribute {
	if statusCode >= 200 && statusCode < 300 {
		return color.FgHiGreen
	}

	if statusCode >= 300 && statusCode < 400 {
		return color.FgHiYellow
	}

	return color.FgHiRed
}

func SplitLines(s string) []string {
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
