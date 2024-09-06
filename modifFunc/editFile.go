package reloadgo

func EditFILE(text [][]string) [][]string {
	input := HexToDecimal(text)
	input = BinToDecimal(input)
	input = PunctModif(input)
	input = Upper(input)
	input = Lower(input)
	// input = PunctModif(input)
	input = Capitalize(input)
	input = SpecialCase(input)
	input = AtoAN(input)
	return input
}
