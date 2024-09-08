package reloadgo

import (
	"strings"
	"unicode"
)

func isPunctuation(char rune) bool {
	return char == '.' || char == ',' || char == '!' || char == '?' || char == ':' || char == ';'
}

func ModifiePunctuation(text string) string {
	var newText string

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
	newText = FixQuote(newText)
	return newText
}

func isSpace(char rune) bool {
	return char == ' '
}

func FixQuote(input string) string {
	lines := strings.Split(input, "\n")
	var finalResult strings.Builder

	for _, line := range lines {
		var result strings.Builder
		inQuotes := false
		startQuote := -1
		for i, ch := range line {
			if ch == '\'' {
				QuoteinsideWord := false
				if i > 0 && !unicode.IsSpace(rune(line[i-1])) && i < len(line)-1 && !unicode.IsSpace(rune(line[i+1])) {
					QuoteinsideWord = true
				}

				if QuoteinsideWord {
					result.WriteRune(ch)
				} else {
					if !inQuotes {
						inQuotes = true
						startQuote = result.Len()
						result.WriteRune(ch)
					} else {
						inQuotes = false
						tempResult := result.String()
						quotedContent := strings.TrimSpace(tempResult[startQuote+1:])
						result.Reset()
						result.WriteString(tempResult[:startQuote+1])
						result.WriteString(quotedContent)
						result.WriteRune('\'')

						if quotedContent == "" && i < len(line)-1 {
							result.WriteRune(' ')
						}
					}
				}
			} else if isSpace(ch) {
				if result.Len() > 0 && result.String()[result.Len()-1] != ' ' {
					result.WriteRune(' ')
				}
			} else {
				result.WriteRune(ch)
			}
		}

		finalResult.WriteString(strings.TrimSpace(result.String()))
		finalResult.WriteRune('\n')
	}

	return strings.TrimSpace(finalResult.String())
}
