package reloadgo

import (
	"strconv"
	"strings"
)

func EditFILE(text [][]string) [][]string {
	// text = PunctModif(text)
	for i := 0; i < len(text); i++ {
		for j := 0; j < len(text[i]); j++ {
			if modifieInstance(text[i][j]) {
				switch text[i][j] {
				case "(hex)":
					modifiedLine := HexToDecimal(text[i][:j])
					modifiedLine = append(modifiedLine, text[i][j+1:]...)
					text[i] = modifiedLine
					j--
				case "(bin)":
					modifiedLine := BinToDecimal(text[i][:j])
					modifiedLine = append(modifiedLine, text[i][j+1:]...)
					text[i] = modifiedLine
					j--
				case "(up)":
					modifiedLine := Upper(text[i][:j])
					modifiedLine = append(modifiedLine, text[i][j+1:]...)
					text[i] = modifiedLine
					j--
				case "(low)":
					modifiedLine := Lower(text[i][:j])
					modifiedLine = append(modifiedLine, text[i][j+1:]...)
					text[i] = modifiedLine
					j--
				case "(cap)":
					modifiedLine := Capitalize(text[i][:j])
					modifiedLine = append(modifiedLine, text[i][j+1:]...)
					text[i] = modifiedLine
					j--
				case "(up,":
					if checkCase(text[i][j+1]) {
						sCase := "up"
						numWord, _ := strconv.Atoi(text[i][j+1][:len(text[i][j+1])-1])
						modifiedLine := SpecialCase(text[i][:j], sCase, numWord)
						modifiedLine = append(modifiedLine, text[i][j+2:]...)
						text[i] = modifiedLine
						j--
					}
				case "(low,":
					if checkCase(text[i][j+1]) {
						sCase := "low"
						numWord, _ := strconv.Atoi(text[i][j+1][:len(text[i][j+1])-1])
						modifiedLine := SpecialCase(text[i][:j], sCase, numWord)
						modifiedLine = append(modifiedLine, text[i][j+2:]...)
						text[i] = modifiedLine
						j--
					}
				case "(cap,":
					if checkCase(text[i][j+1]) {
						sCase := "cap"
						numWord, _ := strconv.Atoi(text[i][j+1][:len(text[i][j+1])-1])
						modifiedLine := SpecialCase(text[i][:j], sCase, numWord)
						modifiedLine = append(modifiedLine, text[i][j+2:]...)
						text[i] = modifiedLine
						j--
					}
				}
			}
		}
		// text[i] = specialMark2(text[i])
	}
	text = PunctModif(text)
	text = AtoAN(text)

	return text
}

func modifieInstance(word string) bool {
	if word == "(hex)" || word == "(bin)" || word == "(up)" || word == "(low)" || word == "(cap)" || word == "(up," || word == "(low," || word == "(cap," {
		return true
	}
	return false
}

func checkCase(text string) bool {
	if strings.HasSuffix(text, ")") {
		if len(text) <= 1 {
			return false
		}
		for i := 0; i < len(text)-1; i++ {
			if text[i] < '0' || text[i] > '9' {
				return false
			}
		}

		return true
	}
	return false
}

// func EditFILE(text [][]string) [][]string {
// 	// input := HexToDecimal(text)
// 	// input = BinToDecimal(input)
// 	// input = PunctModif(input)
// 	// input = Upper(input)
// 	// input = Lower(input)
// 	// // input = PunctModif(input)
// 	// input = Capitalize(input)
// 	// input = SpecialCase(input)
// 	// input = AtoAN(input)
// 	// return input
// }
