package reloadgo

import (
	"strconv"
	"strings"
)

func EditFILE(text [][]string) [][]string {
	// text = PunctModif(text)
	for i := 0; i < len(text); i++ {
		for j := 0; j < len(text[i]); j++ {
			if isModifier(text[i][j]) {
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
	}

	// text = formatText(text)
	text = AtoAN(text)
	//text = FormatPunctuation(text)
	return text
}

func isModifier(word string) bool {
	switch word {
	case "(hex)", "(bin)", "(up)", "(low)", "(cap)", "(up,", "(low,", "(cap,":
		return true
	}
	return false
}

func checkCase(text string) bool {
	if strings.HasSuffix(text, ")") && len(text) > 1 {
		for _, char := range text[:len(text)-1] {
			if char < '0' || char > '9' {
				return false
			}
		}
		return true
	}
	return false
}
