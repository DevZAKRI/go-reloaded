package reloadgo

import (
	"strconv"
)

func HexToDecimal(text []string) []string {
	if len(text) < 1 {
		return []string{}
	} else {
		valDec, _ := strconv.ParseInt(text[len(text)-1], 16, 64)
		text[len(text)-1] = strconv.Itoa(int(valDec))
	}
	return text
}

func BinToDecimal(text []string) []string {
	if len(text) <= 1 {
		return []string{}
	} else {
		valDec, _ := strconv.ParseInt(text[len(text)-1], 2, 64)
		text[len(text)-1] = strconv.Itoa(int(valDec))
	}
	return text
}
