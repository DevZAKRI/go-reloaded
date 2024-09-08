package reloadgo

import (
	"regexp"
	"strings"
)

func isPunctuation(char rune) bool {
	return char == '.' || char == ',' || char == '!' || char == '?' || char == ':' || char == ';'
}

func ModifiePunctuation(text string) string {
	var newText string
	// inQuote := false

	for idx, char := range text {
		var prevChar, nextChar rune
		if idx > 0 {
			prevChar = rune(text[idx-1])
		}
		if idx+1 < len(text) {
			nextChar = rune(text[idx+1])
		}

		if isPunctuation(char) {
			newText += string(char)
			if idx+1 < len(text) && !isPunctuation(rune(text[idx+1])) && !isSpace(rune(text[idx+1])) {
				newText += " "
			}
			continue
		}

		if isSpace(char) {
			if isPunctuation(prevChar) && !isPunctuation(nextChar) {
				newText += " "
				continue
			}
			if !isPunctuation(prevChar) && !isPunctuation(nextChar) {
				newText += " "
				continue
			}
			continue
		}
		newText += string(char)
	}
	newText = formatPunctuation(newText)
	return newText
}

func isSpace(char rune) bool {
	return char == ' '
}

func isApostrophe(char rune) bool {
	return char == '\''
}

func checkApostrophe(text string) bool {
	for _, char := range text {
		if isApostrophe(char) {
			return true
		}
	}
	return false
}

func formatPunctuation(input string) string {
	quotedText := regexp.MustCompile(`'([^']*)'`)
	text := quotedText.ReplaceAllStringFunc(input, func(match string) string {
		content := strings.TrimSpace(match[1 : len(match)-1])
		return "'" + content + "' "
	})
	text = strings.TrimSpace(text)
	nSpaceRemover := regexp.MustCompile(`\s{2,}`)
	return nSpaceRemover.ReplaceAllString(text, " ")
}
