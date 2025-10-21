package utils

import "github.com/fatih/color"

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
