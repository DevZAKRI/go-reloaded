package reloadgo

import "strconv"

func HexChange(text []string) []string {
	newValTExt := ""
	changeDone := false
	var newText []string
	for index, val := range text {
		if val == "(hex)" {
			continue
		}
		if index < len(text)-1 && text[index+1] == "(hex)" {
			newVal, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				panic(err)
			} else {
				newValTExt = strconv.Itoa(int(newVal))
				changeDone = true
			}
		}
		if changeDone {
			newText = append(newText, newValTExt)
		} else {
			newText = append(newText, val)
		}

	}
	return newText
}
