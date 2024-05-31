package main

import (
	"fmt"
)

func main() {
	var week1 Schedule
	week1.initMap()
	//week1.printMap()
	fmt.Printf("To find daily schedule, please, input day of the week: ")
	var findDaySchedule string
	fmt.Scan(&findDaySchedule)
	week1.Search(findDaySchedule)
}

type Schedule struct {
	week1Schedule map[string]string
}

func (s *Schedule) initMap() map[string]string {
	s.week1Schedule = map[string]string{
		"Monday":    "\n1.English;\n2.Mathematics;\n3.Music;\n4.Tennis;",
		"Tuesday":   "\n1.Ckooking;\n2.Sport;\n3.Drawing",
		"Wednesday": "\n1.Mathematics;\n2.Tennis;\n3.English",
		"Thursday":  "\n1.Literature;\n2.Sport;\n3.Reading;\n4.Building",
		"Friday":    "\n1.English;\n2.Sport;",
	}
	return s.week1Schedule
}

func (s *Schedule) printMap() {
	fmt.Println("All elements", s.week1Schedule)
}

func (s *Schedule) Search(key string) {
	if val, ok := s.week1Schedule[key]; ok {
		fmt.Printf("%v schedule is a %v\n", key, val)
	} else {
		fmt.Printf("No such key exists. Try again.\n")
	}
}
