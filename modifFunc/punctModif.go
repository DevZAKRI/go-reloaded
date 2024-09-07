package reloadgo

import "strings"

func PunctModif(text [][]string) [][]string {
	for i := 0; i < len(text); i++ {
		// text[i] = specialMark(text[i])
		for j := 0; j < len(text[i]); j++ {
			word := ""
			pMark := 0
			editLine := text[i][:j]
			for idx, char := range text[i][j] {
				if j > 0 && idx == pMark && isPunctuation(char) {
					text[i][j-1] += string(char)
					pMark++
				} else if len(text[i]) != 1 && len(text[i][j]) == 1 && isPunctuation(char) {
					text[i][j-1] += string(char)
				} else {
					word += string(char)
				}
			}
			editLine = append(editLine, word)
			editLine = append(editLine, text[i][j+1:]...)
			text[i] = editLine
		}
	}
	return text
}

func isPunctuation(char rune) bool {
	switch char {
	case '.':
		return true
	case ',':
		return true
	case '!':
		return true
	case '?':
		return true
	case ':':
		return true
	case ';':
		return true
	default:
		return false
	}
}

func specialMark(text []string) []string {
	inQuate := false
	result := []string{}
	temp := ""

	for i := 0; i < len(text); i++ {
		// If the current string does not contain only a single quote
		if text[i] != "'" {
			for idx := range text[i] {
				// Check for single quotes inside the string
				if text[i][idx] == '\'' {
					if inQuate {
						// Closing quote found inside string, add the trimmed temp
						result = append(result, strings.TrimSpace(temp+"'"))
						temp = ""
						inQuate = false
					} else {
						// Opening quote found, start collecting quoted text
						temp = "'"
						inQuate = true
					}
				} else {
					// Append the character to the temp string inside a quote
					if inQuate {
						temp += string(text[i][idx])
					} else {
						result = append(result, string(text[i][idx]))
					}
				}
			}
		} else {
			// Handle if the current element is a single quote
			if inQuate {
				// Closing quote
				result = append(result, strings.TrimSpace(temp+"'"))
				temp = ""
				inQuate = false
			} else {
				// Opening quote
				temp = "'"
				inQuate = true
			}
		}
	}

	// Append any remaining text
	if temp != "" {
		result = append(result, strings.TrimSpace(temp))
	}

	return result
}

// func specialMark2(text []string) []string {
// 	inQuote := false
// 	var result []string
// 	for i := 0; i < len(text); i++ {
// 		if text[i] == "'" && !inQuote && i < len(text)-1 {
// 			text[i] += text[i+1]
// 			result = append(result, text[:i+1]...)
// 			if i < len(text)-2 {
// 				result = append(result, text[i+2:]...)
// 			}
// 		} else if text[i] == "'" && inQuote && i < len(text)-1 && i > 0 {
// 			text[i-1] += text[i]
// 			result = append(result, text[:i]...)
// 			if i < len(text)-1 {
// 				result = append(result, text[i+1:]...)
// 			}
// 		}
// 	}
// 	// Return the modified slice of strings
// 	return result
// }
