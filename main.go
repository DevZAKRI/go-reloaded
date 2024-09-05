package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <input_file> <output_file names")
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

	content, err := os.ReadFile(file.Name())
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(content)
	fileContent := strings.Fields(string(content))
	newFileContent := []string{}

	for i := 0; i < len(fileContent); i++ {
		val := fileContent[i]
		if i < len(fileContent)-1 && fileContent[i+1] == "(hex)" {
			valDec, err := strconv.ParseInt(val, 16, 64)
			if err != nil {
				fmt.Println("Error converting hex to decimal:", err)
				return
			}
			newFileContent = append(newFileContent, strconv.Itoa(int(valDec)))
			i++
		} else {
			newFileContent = append(newFileContent, val)
		}
	}

	err = os.WriteFile(outputFile, []byte(strings.Join(newFileContent, " ")), 0o644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Processing complete. Output written to", outputFile)
}
