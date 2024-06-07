package index

import (
	"strings"
)

type Index struct {
	index map[string][]int
	text  []string
}

func New(text []string) *Index {
	idx := &Index{
		index: make(map[string][]int),
		text:  text,
	}
	for i, line := range text {
		words := strings.Fields(line)
		for _, word := range words {
			word = strings.ToLower(word)
			idx.index[word] = append(idx.index[word], i)
		}
	}
	return idx
}

type SearchResult struct {
	SearchWord  string
	LineNumbers []int
}

func (i *Index) Search(word string) SearchResult {
	//var inputSearchKey string
	//fmt.Println("Input word to search")
	//fmt.Scan(&inputSearchKey)
	//inputSearchKey = strings.ToLower(inputSearchKey)
	word = strings.ToLower(word)
	if val, ok := i.index[word]; ok {
		//fmt.Printf("The word '%v' exists in lines# %v\n", word, val)
		return SearchResult{
			SearchWord:  word,
			LineNumbers: val,
		}
	}
	return SearchResult{
		SearchWord:  word,
		LineNumbers: []int{},
	}

}
