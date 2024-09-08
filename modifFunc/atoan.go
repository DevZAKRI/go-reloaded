package reloadgo

func AtoAN(text [][]string) [][]string {
	for i := range text {
		for j := range text[i] {
			if isA(text[i][j]) && j < len(text[i])-1 && isVowel(string(text[i][j+1][0])) {
				text[i][j] = replaceA(text[i][j])
			}
		}
	}
	return text
}

func isA(word string) bool {
	return word == "a" || word == "A"
}

func replaceA(word string) string {
	if word == "a" {
		return "an"
	}
	return "An"
}

func isVowel(char string) bool {
	switch char {
	case "a", "A", "e", "E", "i", "I", "o", "O", "u", "U", "h", "H":
		return true
	default:
		return false
	}
}
