package main

import (
	"Homework5/index"
	"bufio"
	"fmt"
	"os"
)

func main() {
	text := chooseTextOrigin()
	index := index.New(text)
	var word string
	fmt.Println("Enter a word to search:")
	fmt.Scan(&word)
	searchResult := index.Search(word)
	fmt.Printf("Search results for '%v': %v\n", searchResult.SearchWord, searchResult.LineNumbers)
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

func getTextFromTerminal() []string {
	fmt.Printf("How manu lines are you going to input?  ")
	var number int
	fmt.Scan(&number)
	fmt.Scanln()
	text := make([]string, number)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < number; i++ {
		fmt.Printf("Input line #%v  -  ", i)
		if scanner.Scan() {
			text[i] = scanner.Text()
		} else {
			if err := scanner.Err(); err != nil {
				fmt.Println("Error reading input:", err)
			}
		}
	}
	return text
}

func chooseTextOrigin() []string {
	fmt.Println("Input 1 or 2 if you want to get text from: '1' - file; '2' - terminal")
	var textOrigin int
	fmt.Scan(&textOrigin)
	var text []string
	var finish bool
	for !finish {
		switch textOrigin {
		case 1:
			text = getTextFromFile()
			finish = true
		case 2:
			text = getTextFromTerminal()
			finish = true
		default:
			fmt.Println("Invalid choice. Please choose a valid number.")
			fmt.Scan(&textOrigin)
		}
	}
	return text
}
