package reloadgo

func isPunctuation(char rune) bool {
	return char == '.' || char == ',' || char == '!' || char == '?' || char == ':' || char == ';'
}

func PunctModif(text [][]string) [][]string {
	for i := 0; i < len(text); i++ {
		for j := 0; j < len(text[i]); j++ {
			word := ""
			pMark := 0
			editLine := text[i][:j]
			for idx, char := range text[i][j] {
				if j > 0 && idx == pMark && isPunctuation(char) {
					text[i][j-1] += string(char)
					pMark++
				} else if j > 0 && len(text[i]) != 1 && len(text[i][j]) == 1 && isPunctuation(char) {
					text[i][j-1] += string(char)
				} else {
					word += string(char)
				}
			}
			editLine = append(editLine, word)
			if j+1 < len(text[i]) { // Check to prevent out of range
				editLine = append(editLine, text[i][j+1:]...)
			}
			text[i] = editLine
		}
	}
	return text
}

// func PunctModif(text [][]string) [][]string {
// 	for i := 0; i < len(text); i++ {
// 		// text[i] = specialMark(text[i])
// 		for j := 0; j < len(text[i]); j++ {
// 			word := ""
// 			pMark := 0
// 			editLine := text[i][:j]
// 			for idx, char := range text[i][j] {
// 				if j > 0 && idx == pMark && isPunctuation(char) {
// 					text[i][j-1] += string(char)
// 					pMark++
// 				} else if len(text[i]) != 1 && len(text[i][j]) == 1 && isPunctuation(char) {
// 					text[i][j-1] += string(char)
// 				} else {
// 					word += string(char)
// 				}
// 			}
// 			editLine = append(editLine, word)
// 			editLine = append(editLine, text[i][j+1:]...)
// 			text[i] = editLine
// 		}
// 	}
// 	return text
// }
