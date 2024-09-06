package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	reloadgo "reloadgo/modifFunc"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("u need to have two argument \"input file name\" and \"output file name\"")
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var content []string
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	fileContent := make([][]string, len(content))
	for i := 0; i < len(content); i++ {
		fileContent[i] = append(fileContent[i], strings.Fields(content[i])...)
	}
	text := reloadgo.EditFILE(fileContent)
	// fmt.Printf("%v", text[1][3][:len(text[1][3])-1])
	fmt.Println(text)
	// for i := 0; i < len(text); i++ {
	// 	err = os.WriteFile(outputFile, []byte(strings.Join(text[i], " ")), 0o644)
	// 	if err != nil {
	// 		fmt.Println("Error writing to file:", err)
	// 		return
	// 	}
	// }
	outputText := ""
	for i := 0; i < len(text); i++ {
		for j := 0; j < len(text[i]); j++ {
			if j < len(text[i])-1 {
				outputText += text[i][j] + " "
			} else {
				if i < len(text)-1 {
					outputText += text[i][j] + "\n"
				} else {
					outputText += text[i][j]
				}
			}
		}
	}
	fmt.Printf("%v", outputText)
	err = os.WriteFile(outputFile, []byte(outputText), 0o777)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
