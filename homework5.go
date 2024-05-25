package main

import (
	"fmt"
)

// чи можна виковистовувати методи для мапів окремо, не вкладаючи їх в структуру?
func main() {

	var spaceObjects SpaceObjects

	spaceObjects.inputMap()
	spaceObjects.printMap()

	fmt.Printf("Input Key: ")
	var findObject string
	fmt.Scan(&findObject)

	spaceObjects.Search(findObject)
}

type SpaceObjects struct {
	objects map[string]string
}

func (s *SpaceObjects) inputMap() {
	s.objects = map[string]string{
		"Sun":   "star",
		"Mars":  "planet",
		"Moon":  "satellite",
		"Earth": "planet",
	}
}

func (s *SpaceObjects) printMap() {
	fmt.Println("All elements of map", s.objects)
}

// Example:
//var people = map[string]int{
//    "Tom": 1,
//    "Bob": 2,
//    "Sam": 4,
//    "Alice": 8,
// }
//if val, ok := people["Tom"]; ok{
//    fmt.Println(val)
//}

func (s *SpaceObjects) Search(key string) {
	if why, yes := s.objects[key]; yes {
		fmt.Printf("%v is a %v\n", key, why)
	} else {
		fmt.Printf("No such key exists. Try again.\n")
	}
}
