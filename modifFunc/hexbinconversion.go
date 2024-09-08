package reloadgo

import (
	"strconv"
)

func HexToDecimal(text []string) []string {
	if len(text) == 0 {
		return []string{}
	}
	valDec, err := strconv.ParseInt(text[len(text)-1], 16, 64)
	if err == nil {
		text[len(text)-1] = strconv.Itoa(int(valDec))
	}
	return text
}

func BinToDecimal(text []string) []string {
	if len(text) == 0 {
		return []string{}
	}
	valDec, err := strconv.ParseInt(text[len(text)-1], 2, 64)
	if err == nil {
		text[len(text)-1] = strconv.Itoa(int(valDec))
	}
	return text
}
