package main

import (
	"fmt"
	"sort"
)

type person struct {
	name   string
	age    int
	height int
}

func main() {
	personSlice := make([]person, 0, 3)
	//inits personSlice to a slice that has a capacity of three, but nothing is filled yet

	// fmt.Printf("value in personSlice[0] %v\n", personSlice[0])
	/*
			panic: runtime error: index out of range [0] with length 0

				goroutine 1 [running]:
				main.main()
			/home/nindgabeet/workspace/github.com/sohWenMing/mastering_go/slices/sorting_slices/sorting_slices.go:15 +0x12
		exit status 2
	*/
	personSlice = append(personSlice, person{"wen", 42, 167})
	personSlice = append(personSlice, person{"sarah", 39, 150})
	personSlice = append(personSlice, person{"hailey", 7, 100})
	fmt.Println("personSlice before sorting: ", personSlice)

	sort.Slice(personSlice, func(i, j int) bool {
		return personSlice[i].age < personSlice[j].age
	})
	fmt.Println("personSlice after sorting: ", personSlice)
}
