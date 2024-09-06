package reloadgo

func AtoAN(text [][]string) [][]string {
	for i := 0; i < len(text); i++ {
		for j := 0; j < len(text[i]); j++ {
			if text[i][j] == "a" && j < len(text[i])-1 {
				if isVowel(string(text[i][j+1][0])) {
					text[i][j] = "an"
				}
			} else if text[i][j] == "A" && j < len(text[i])-1 {
				if isVowel(string(text[i][j+1][0])) {
					text[i][j] = "AN"
				}
			}
		}
	}
	return text
}

func isVowel(char string) bool {
	switch char {
	case "a", "A":
		return true
	case "e", "E":
		return true
	case "i", "I":
		return true
	case "u", "U":
		return true
	case "o", "O":
		return true
	case "h", "H":
		return true
	default:
		return false
	}
}
