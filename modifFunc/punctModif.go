package reloadgo

func isPunctuation(char rune) bool {
	return char == '.' || char == ',' || char == '!' || char == '?' || char == ':' || char == ';'
}

func ModifiePunctuation(text string) string {
	newText := ""
	for idx, char := range text {
		if isPunctuation(char) {
			if idx == 0 || idx == len(text)-1 {
				newText += string(char)
				if idx == 0 && len(text) > 1 && !isPunctuation(rune(text[idx-1])) {
					newText += " "
				}
			} else if isSpace(text[idx-1]) && idx < len(text)-1 && !isPunctuation(rune(text[idx-1])) {
				newText += " " + string(char)
			}
		} else {
			newText += string(char)
		}
	}
	return newText
}

func isSpace(char byte) bool {
	return char == ' '
}
