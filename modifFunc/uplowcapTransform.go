package reloadgo

import (
	"strings"
)

func Upper(text []string) []string {
	if len(text) == 0 {
		return text
	}
	text[len(text)-1] = strings.ToUpper(text[len(text)-1])
	return text
}

func Lower(text []string) []string {
	if len(text) == 0 {
		return text
	}
	text[len(text)-1] = strings.ToLower(text[len(text)-1])
	return text
}

func Capitalize(text []string) []string {
	if len(text) == 0 {
		return text
	}
	text[len(text)-1] = CapitalizeWord2(strings.ToLower(text[len(text)-1]))
	return text
}

func SpecialCase(text []string, sCase string, nWords int) []string {
	if len(text) == 0 {
		return text
	}
	if nWords > len(text) {
		nWords = len(text)
	}
	for k := 0; k < nWords; k++ {
		switch sCase {
		case "up":
			text[len(text)-1-k] = strings.ToUpper(text[len(text)-1-k])
		case "low":
			text[len(text)-1-k] = strings.ToLower(text[len(text)-1-k])
		case "cap":
			text[len(text)-1-k] = CapitalizeWord2(text[len(text)-1-k])
		}
	}
	return text
}

// capitalizes the first letter of a word. ignore modification if first letter is a special character
func CapitalizeWord(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + word[1:]
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

// capitalize first letter in a word ignore special char
func CapitalizeWord2(word string) string {
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
