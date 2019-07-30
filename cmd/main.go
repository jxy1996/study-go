package main

import (
	"fmt"
	"github.com/jxy1996/study-go/sorter"
	"sort"
)

func main() {
	p := sorter.Persons{
		sorter.Person{FirstName: "a", LastName: "b"},
		sorter.Person{FirstName: "b", LastName: "c"},
		sorter.Person{FirstName: "c", LastName: "d"},
	}

	fmt.Println(p)
	sort.Sort(p)
	fmt.Println(p)
}
