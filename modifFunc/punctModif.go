package reloadgo

func PunctModif(text [][]string) [][]string {
	for i := 0; i < len(text); i++ {
		for j := 0; j < len(text[i]); j++ {
			word := ""
			pMark := 0
			editLine := text[i][:j]
			for idx, char := range text[i][j] {
				if len(text[i][j]) > 1 && idx == pMark && isPunctuation(char) {
					text[i][j-1] += string(char)
					pMark++
				} else if len(text[i][j]) == 1 && isPunctuation(char) {
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

func specialcas(){

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
