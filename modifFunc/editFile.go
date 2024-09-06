package reloadgo

func EditFILE(text [][]string) [][]string {
	text = PunctModif(text)
	for i := 0; i < len(text); i++ {
		for j := 0; j < len(text[i]); j++ {
			if modifieInstance(text[i][j]) {
				switch text[i][j] {
				case "(hex)":
					modifiedLine := HexToDecimal(text[i][:j])
					modifiedLine = append(modifiedLine, text[i][j+1:]...)
					text[i] = modifiedLine
				case "(bin)":
					modifiedLine := BinToDecimal(text[i][:j])
					modifiedLine = append(modifiedLine, text[i][j+1:]...)
					text[i] = modifiedLine
				case "(up)":
					modifiedLine := Upper(text[i][:j])
					modifiedLine = append(modifiedLine, text[i][j+1:]...)
					text[i] = modifiedLine
				case "(low)":
					modifiedLine := Lower(text[i][:j])
					modifiedLine = append(modifiedLine, text[i][j+1:]...)
					text[i] = modifiedLine
				case "(cap)":
					modifiedLine := Capitalize(text[i][:j])
					modifiedLine = append(modifiedLine, text[i][j+1:]...)
					text[i] = modifiedLine
				case "(up,":
				case "(low,":
				case "(cap,":
				}
			}
		}
	}
	return text
}

func modifieInstance(word string) bool {
	if word == "(hex)" || word == "(bin)" || word == "(up)" || word == "(low)" || word == "(cap)" || word == "(up," || word == "(low," || word == "(cap," {
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
