package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func countBytes(text string) int {
	return len([]byte(strings.TrimSpace(text)))
}

func countWords(text string) int {
	return len(strings.Fields(text))
}

func countLines(text string) int {
	return len(strings.Split(text, "\n"))
}

func countChars(text string) int {
	return utf8.RuneCountInString(text)
}

func readFromInput(option string) string {

	scanner := bufio.NewScanner(os.Stdin)

	var lines string

	for scanner.Scan() {

		line := scanner.Text()

		lines = lines + "\n" + line
	}

	fmt.Println(lines)

	return lines

}

func main() {

	var option, fileName, text string
	args := os.Args

	if len(args) < 1 {
		log.Fatal("Incorrect command line arguments")
		os.Exit(1)
	}

	if args[1][0] == '-' {
		option = args[1]
		if len(args) > 2 {
			fileName = args[2]
		}
	} else {
		option = "default"
		fileName = args[1]
	}

	fileContent, err := os.ReadFile(fileName)

	if err != nil {
		fileContent = []byte(readFromInput(option))
	}

	text = string(fileContent)

	if option == "-c" {
		fmt.Printf("%v %s\n", countBytes(text), fileName)
	} else if option == "-l" {
		fmt.Printf("%v %s\n", countLines(text), fileName)
	} else if option == "-w" {
		fmt.Printf("%v %s\n", countWords(text), fileName)
	} else if option == "-m" {
		fmt.Printf("%v %s\n", countChars(text), fileName)
	} else {
		fmt.Printf("%v %v %v %v\n", countLines(text), countWords(text), countBytes(text), fileName)
	}

}
