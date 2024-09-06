package reloadgo

import (
	"strconv"
	"strings"
)

func Upper(text []string) []string {
	if len(text) < 1 {
		return []string{}
	} else {
		text[len(text)-1] = strings.ToUpper(text[len(text)-1])
	}
	return text
}

func Lower(text []string) []string {
	if len(text) < 1 {
		return []string{}
	} else {
		text[len(text)-1] = strings.ToLower(text[len(text)-1])
	}
	return text
}

func Capitalize(text []string) []string {
	if len(text) < 1 {
		return []string{}
	} else {
		text[len(text)-1] = capWORD(text[len(text)-1])
	}
	return text
}

// func Upper(text [][]string) [][]string {
// 	nLines := len(text)

// 	for i := 0; i < nLines; i++ {
// 		var newLine []string
// 		for j := 0; j < len(text[i]); j++ {
// 			if j == 0 && text[i][j] == "(up)" {
// 				continue
// 			}
// 			if j < len(text[i])-1 && text[i][j+1] == "(up)" {
// 				newLine = append(newLine, strings.ToUpper(text[i][j]))
// 				j++
// 			} else {
// 				newLine = append(newLine, text[i][j])
// 			}
// 		}
// 		text[i] = newLine
// 	}
// 	return text
// }

// func Lower(text [][]string) [][]string {
// 	nLines := len(text)

// 	for i := 0; i < nLines; i++ {
// 		var newLine []string
// 		for j := 0; j < len(text[i]); j++ {
// 			if j == 0 && text[i][j] == "(low)" {
// 				continue
// 			}
// 			if j < len(text[i])-1 && text[i][j+1] == "(low)" {
// 				newLine = append(newLine, strings.ToLower(text[i][j]))
// 				j++
// 			} else {
// 				newLine = append(newLine, text[i][j])
// 			}
// 		}
// 		text[i] = newLine
// 	}
// 	return text
// }

// func Capitalize(text [][]string) [][]string {
// 	nLines := len(text)

// 	for i := 0; i < nLines; i++ {
// 		var newLine []string
// 		for j := 0; j < len(text[i]); j++ {
// 			if j == 0 && text[i][j] == "(cap)" {
// 				continue
// 			}
// 			if j < len(text[i])-1 && text[i][j+1] == "(cap)" {
// 				word := strings.ToLower(text[i][j])
// 				newWord := ""
// 				for idx, char := range word {
// 					if idx == 0 {
// 						newWord += string(char - 32)
// 					} else {
// 						newWord += string(char)
// 					}
// 				}
// 				newLine = append(newLine, newWord)
// 				j++
// 			} else {
// 				newLine = append(newLine, text[i][j])
// 			}
// 		}
// 		text[i] = newLine
// 	}
// 	return text
// }

func SpecialCase(text [][]string) [][]string {
	nLines := len(text)
	for i := 0; i < nLines; i++ {
		// var newLine []string
		for j := 0; j < len(text[i]); j++ {
			if j < len(text[i])-2 && (text[i][j+1] == "(cap," || text[i][j+1] == "(low," || text[i][j+1] == "(up,") {
				switch text[i][j+1] {
				case "(low,":
					numWord, _ := strconv.Atoi(text[i][j+2][:len(text[i][j+2])-1])
					if numWord > len(text[i][:j+1]) {
						numWord = len(text[i][:j+1])
					}
					for k := 0; k < numWord; k++ {
						text[i][j-k] = strings.ToLower(text[i][j-k])
					}
					firstHalf := text[i][:j+1]
					secondHalf := text[i][j+3:]
					firstHalf = append(firstHalf, secondHalf...)
					text[i] = firstHalf
				case "(up,":
					numWord, _ := strconv.Atoi(text[i][j+2][:len(text[i][j+2])-1])
					if numWord > len(text[i][:j+1]) {
						numWord = len(text[i][:j+1])
					}
					for k := 0; k < numWord; k++ {
						text[i][j-k] = strings.ToUpper(text[i][j-k])
					}
					firstHalf := text[i][:j+1]
					secondHalf := text[i][j+3:]
					firstHalf = append(firstHalf, secondHalf...)
					text[i] = firstHalf
				case "(cap,":
					numWord, _ := strconv.Atoi(text[i][j+2][:len(text[i][j+2])-1])
					if numWord > len(text[i][:j+1]) {
						numWord = len(text[i][:j+1])
					}
					for k := 0; k < numWord; k++ {
						text[i][j-k] = capWORD(text[i][j-k])
					}
					firstHalf := text[i][:j+1]
					secondHalf := text[i][j+3:]
					firstHalf = append(firstHalf, secondHalf...)
					text[i] = firstHalf
				}
			}
		}
		// text[i] = newLine
	}
	return text
}

func capWORD(word string) string {
	newWord := ""
	for idx, char := range word {
		if idx == 0 {
			newWord += string(char - 32)
		} else {
			newWord += string(char)
		}
	}
	return newWord
}
