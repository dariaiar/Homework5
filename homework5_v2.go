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
	//fmt.Println("Your text is:  ", text)
	//for b, rangedText := range text {
	//	fmt.Println(b, rangedText)
	//}
	return text
}

//func newIndex(s index.searchByIndex) (text []string, indexText map[string][]int) {
//	index.searchByIndex(text)(indexText)
//	index.Search(indexText)
//	return
//}

//func insertDataToMap(text []string) map[string][]int {
//	indexText := make(map[string][]int)
//	for i, line := range text {
//		words := strings.Fields(line)
//		for _, word := range words {
//			word = strings.ToLower(word)
//			indexText[word] = append(indexText[word], i)
//
//		}
//	}
//	return indexText
//}

//func search(indexText map[string][]int) {
//	var inputSearchKey string
//	fmt.Println("Input word to search")
//	fmt.Scan(&inputSearchKey)
//	inputSearchKey = strings.ToLower(inputSearchKey)
//	if val, ok := indexText[inputSearchKey]; ok {
//		fmt.Printf("The word '%v' exists in lines# %v\n", inputSearchKey, val)
//	} else {
//		fmt.Printf("No such key exists. Try again.\n")
//	}
//}
