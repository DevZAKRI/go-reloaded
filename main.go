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
		fmt.Println("You need to provide two file names: \"input file name\" and \"output file name\"")
		return
	}

	inputFile, outputFile := os.Args[1], os.Args[2]
	fileContent, err := readFile(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	EditedContent := reloadgo.EditFILE(fileContent)
	outputText := formatOutput(EditedContent)
	finalText := reloadgo.ModifiePunctuation(outputText)
	fmt.Println(finalText)

	if err := writeFile(outputFile, outputText); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func readFile(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var content [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, strings.Fields(scanner.Text()))
	}
	return content, scanner.Err()
}

func formatOutput(content [][]string) string {
	var outputBuilder strings.Builder
	for i, line := range content {
		outputBuilder.WriteString(strings.Join(line, " "))
		if i < len(content)-1 {
			outputBuilder.WriteString("\n")
		}
	}

	return outputBuilder.String()
}

func writeFile(fileName, content string) error {
	return os.WriteFile(fileName, []byte(content), 0o777)
}

// func main() {
// 	if len(os.Args) != 3 {
// 		fmt.Println("u need to have two argument \"input file name\" and \"output file name\"")
// 		return
// 	}
// 	inputFile := os.Args[1]
// 	outputFile := os.Args[2]
// 	file, err := os.Open(inputFile)
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return
// 	}
// 	defer file.Close()
// 	scanner := bufio.NewScanner(file)
// 	var content []string
// 	for scanner.Scan() {
// 		content = append(content, scanner.Text())
// 	}
// 	fileContent := make([][]string, len(content))
// 	for i := 0; i < len(content); i++ {
// 		fileContent[i] = append(fileContent[i], strings.Fields(content[i])...)
// 	}
// 	text := reloadgo.EditFILE(fileContent)
// 	// fmt.Println(text)
// 	outputText := ""
// 	for i := 0; i < len(text); i++ {
// 		for j := 0; j < len(text[i]); j++ {
// 			if j < len(text[i])-1 {
// 				outputText += text[i][j] + " "
// 			} else {
// 				if i < len(text)-1 {
// 					outputText += text[i][j] + "\n"
// 				} else {
// 					outputText += text[i][j]
// 				}
// 			}
// 		}
// 	}
// 	fmt.Printf("%v", outputText)
// 	err = os.WriteFile(outputFile, []byte(outputText), 0o777)
// 	if err != nil {
// 		fmt.Println("Error writing to file:", err)
// 		return
// 	}
// }
