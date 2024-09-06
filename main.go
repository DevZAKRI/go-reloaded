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
		fmt.Println("Usage: go run main.go <input_file> <output_file names")
		return
	}
	inputFile := os.Args[1]
	// outputFile := os.Args[2]
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
	// err = os.WriteFile(outputFile, []byte(strings.Join(newFileContent, " ")), 0o644)
	// if err != nil {
	// 	fmt.Println("Error writing to file:", err)
	// 	return
	// }
}
