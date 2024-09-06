package reloadgo

import (
	"strconv"
)

func HexToDecimal(text [][]string) [][]string {
	nLines := len(text)

	for i := 0; i < nLines; i++ {
		var newLine []string
		for j := 0; j < len(text[i]); j++ {
			if j < len(text[i])-1 && text[i][j+1] == "(hex)" {
				valDec, _ := strconv.ParseInt(text[i][j], 16, 64)
				newLine = append(newLine, strconv.Itoa(int(valDec)))
				j++
			} else {
				newLine = append(newLine, text[i][j])
			}
		}
		text[i] = newLine
	}
	return text
}

func BinToDecimal(text [][]string) [][]string {
	nLines := len(text)

	for i := 0; i < nLines; i++ {
		var newLine []string
		for j := 0; j < len(text[i]); j++ {
			if j < len(text[i])-1 && text[i][j+1] == "(bin)" {
				valDec, _ := strconv.ParseInt(text[i][j], 2, 64)
				newLine = append(newLine, strconv.Itoa(int(valDec)))
				j++
			} else {
				newLine = append(newLine, text[i][j])
			}
		}
		text[i] = newLine
	}
	return text
}
