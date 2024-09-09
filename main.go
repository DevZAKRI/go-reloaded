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
		fmt.Println("Encounter the following error: \"", err, "\" fix and try again:")
		return
	}
	if !strings.HasSuffix(outputFile, ".txt") && !strings.HasSuffix(inputFile, ".txt") {
		fmt.Println("The output file SHOULD NOT be a go file")
		return
	}
	finalText := reloadgo.EditFILE(fileContent)

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

func writeFile(fileName, content string) error {
	return os.WriteFile(fileName, []byte(content), 0o777)
}
