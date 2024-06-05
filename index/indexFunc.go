package index

import (
	"fmt"
	"strings"
)

type Index struct{}

func (s Index) NewIndex(text []string) map[string][]int {
	indexText := make(map[string][]int)
	for i, line := range text {
		words := strings.Fields(line)
		for _, word := range words {
			word = strings.ToLower(word)
			indexText[word] = append(indexText[word], i)

		}
	}
	return indexText
}

func (s Index) Search(indexText map[string][]int) {
	var inputSearchKey string
	fmt.Println("Input word to search")
	fmt.Scan(&inputSearchKey)
	inputSearchKey = strings.ToLower(inputSearchKey)
	if val, ok := indexText[inputSearchKey]; ok {
		fmt.Printf("The word '%v' exists in lines# %v\n", inputSearchKey, val)
	} else {
		fmt.Printf("No such key exists. Try again.\n")
	}
}
