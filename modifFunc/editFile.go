package reloadgo

func EditFILE(text [][]string) [][]string {
	input := HexToDecimal(text)
	input = BinToDecimal(input)
	input = Upper(input)
	input = Lower(input)
	input = Capitalize(input)
	input = SpecialCase(input)
	return input
}
