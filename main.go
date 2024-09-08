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

	if err := writeFile(outputFile, finalText); err != nil {
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
