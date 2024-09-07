package reloadgo

func isPunctuation(char rune) bool {
	return char == '.' || char == ',' || char == '!' || char == '?' || char == ':' || char == ';'
}

func ModifiePunctuation(text string) string {
	newText := ""
	inQuote := false
	for idx, char := range text {
		if isPunctuation(char) {
			// Punc = true
			newText += string(char)
			continue
		}
		if isSpace(byte(char)) {
			if isPunctuation(rune(text[idx-1])) && !isPunctuation(rune(text[idx+1])) {
				newText += " "
				continue
			} else if !isPunctuation(rune(text[idx-1])) && !isPunctuation(rune(text[idx+1])) {
				newText += " "
				continue
			} else {
				continue
			}
		}
		newText += string(char)
	}
	return newText
}

func isSpace(char byte) bool {
	return char == ' '
}

func isApostrophe(char rune) bool {
	return char == '\''
}
