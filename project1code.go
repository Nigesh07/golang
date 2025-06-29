package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	ifile, err := os.Open("project1.go.txt")
	if err != nil {
		fmt.Println("Error1")
	}
	ofile, err := os.Create("error")
	if err != nil {
		fmt.Println("Error2")
	}

	scanner := bufio.NewScanner(ifile)
	writer := bufio.NewWriter(ofile)
	word := "ERROR"
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			_, err := writer.WriteString(line + "\n")
			if err != nil {
				fmt.Println("Error3")
			}
		}

	}
	writer.Flush()
}
