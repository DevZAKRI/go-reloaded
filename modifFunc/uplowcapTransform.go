package reloadgo

import (
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
	// newWord := ""
	if len(text) < 1 {
		return []string{}
	} else {
		word := strings.ToLower(text[len(text)-1])
		newWord := CapitalizeWord(word)
		text[len(text)-1] = newWord
	}
	return text
}

func SpecialCase(text []string, sCase string, nWords int) []string {
	if len(text) < 1 {
		return []string{}
	} else {
		if nWords > len(text) {
			nWords = len(text)
		}
		switch sCase {
		case "up":
			for k := 0; k < nWords; k++ {
				text[len(text)-1-k] = strings.ToUpper(text[len(text)-1-k])
			}
		case "low":
			for k := 0; k < nWords; k++ {
				text[len(text)-1-k] = strings.ToLower(text[len(text)-1-k])
				// fmt.Println(text[len(text)-1-k])
			}
		case "cap":
			for k := 0; k < nWords; k++ {
				text[len(text)-1-k] = CapitalizeWord(text[len(text)-1-k])
			}
		}
	}
	return text
}

// func SpecialCase(text [][]string) [][]string {
// 	nLines := len(text)
// 	for i := 0; i < nLines; i++ {
// 		// var newLine []string
// 		for j := 0; j < len(text[i]); j++ {
// 			if j < len(text[i])-2 && (text[i][j+1] == "(cap," || text[i][j+1] == "(low," || text[i][j+1] == "(up,") {
// 				switch text[i][j+1] {
// 				case "(low,":
// 					numWord, _ := strconv.Atoi(text[i][j+2][:len(text[i][j+2])-1])
// 					if numWord > len(text[i][:j+1]) {
// 						numWord = len(text[i][:j+1])
// 					}
// 					for k := 0; k < numWord; k++ {
// 						text[i][j-k] = strings.ToLower(text[i][j-k])
// 					}
// 					firstHalf := text[i][:j+1]
// 					secondHalf := text[i][j+3:]
// 					firstHalf = append(firstHalf, secondHalf...)
// 					text[i] = firstHalf
// 				case "(up,":
// 					numWord, _ := strconv.Atoi(text[i][j+2][:len(text[i][j+2])-1])
// 					if numWord > len(text[i][:j+1]) {
// 						numWord = len(text[i][:j+1])
// 					}
// 					for k := 0; k < numWord; k++ {
// 						text[i][j-k] = strings.ToUpper(text[i][j-k])
// 					}
// 					firstHalf := text[i][:j+1]
// 					secondHalf := text[i][j+3:]
// 					firstHalf = append(firstHalf, secondHalf...)
// 					text[i] = firstHalf
// 				case "(cap,":
// 					numWord, _ := strconv.Atoi(text[i][j+2][:len(text[i][j+2])-1])
// 					if numWord > len(text[i][:j+1]) {
// 						numWord = len(text[i][:j+1])
// 					}
// 					for k := 0; k < numWord; k++ {
// 						text[i][j-k] = capWORD(text[i][j-k])
// 					}
// 					firstHalf := text[i][:j+1]
// 					secondHalf := text[i][j+3:]
// 					firstHalf = append(firstHalf, secondHalf...)
// 					text[i] = firstHalf
// 				}
// 			}
// 		}
// 		// text[i] = newLine
// 	}
// 	return text
// }

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

func CapitalizeWord(word string) string {
	newWord := ""
	for i := 0; i < len(word); i++ {
		if isPunctuation(rune(word[i])) || word[i] == '\'' {
			newWord += string(word[i])
		} else if !(word[i] >= 'a' && word[i] <= 'z') {
			newWord += string(word[i])
		} else {
			newWord += capWORD(string(word[i])) + word[i+1:]
			break
		}
	}
	return newWord
}
