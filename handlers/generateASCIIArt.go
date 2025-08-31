package ourcode

import (
	"fmt"
	"os"
	"strings"
)

func GenerateASCIIArt(text, banner string) (string, error) {
	bannerFile := fmt.Sprintf("banners/%s.txt", banner)

	content, err := os.ReadFile(bannerFile)
	if err != nil {
		return "", fmt.Errorf("could not read banner file: %v", err)
	}

	lines := strings.Split(string(content), "\n")
	if banner == "thinkertoy" || banner == "shadow" {
		lines = strings.Split(string(content), "\r\n")
	}

	textLines := strings.Split(text, "\r\n")

	result := ""
	for _, line := range textLines {
		if line == "" {
			result += "\n"
			continue
		}

		for i := 1; i < 9; i++ {
			for _, char := range line {
				if char < 32 || char > 126 {
					result = "You don't put a caracter from ascii 😆"
					break
				}

				startLine := int(char-32)*9 + i
				if startLine < len(lines) && startLine >= 0 {
					result += lines[startLine]
				}
			}
			result += "\n"
		}
	}

	return result, nil
}
