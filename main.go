package main

import (
	"fmt"
	"os"

	reloadgo "reloadgo/modFunc"
)

func main() {
	// if len(os.Args[1:]) < 2 {
	// 	fmt.Println("Missing argument!!")
	// 	return
	// } else if len(os.Args[1:]) > 2 {
	// 	fmt.Println("You have more argument than required")
	// 	return
	// }
	inputFile := os.Args[1]
	// outputFile := os.Args[2]
	// var ArrangedText [][]string
	var fileContent []string
	word := ""
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("no such a file exist")
	} else {
		content, err := os.ReadFile(file.Name())
		if err != nil {
			fmt.Println("error reading file")
		} else {
			for _, val := range content {
				if val == ' ' {
					fileContent = append(fileContent, word)
					word = ""
				}
				word += string(val)

			}
			fileContent = append(fileContent, word)
			file.Close()
		}
		afterHex := reloadgo.HexChange(fileContent)
		fmt.Println(afterHex)
	}
}
