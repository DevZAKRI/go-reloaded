package reloadgo

func EditFILE(text [][]string) [][]string {
	input := HexToDecimal(text)
	input = BinToDecimal(input)
	return input
}
