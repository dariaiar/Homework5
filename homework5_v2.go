package main

import (
	"Homework5/index"
	"bufio"
	"fmt"
	"os"
)

func main() {
	text := getTextFromFile()
	indexText := index.NewIndex(text)
	res := index.Search(indexText)

	fmt.Println(res)

}

func getTextFromFile() []string {
	myFile, err := os.Open("yellowSubmarine.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer myFile.Close()
	var text []string
	scanner := bufio.NewScanner(myFile)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		text = append(text, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}

	return text
}
