package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{5, 6, 7, 8}

	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
	/*
		[1, 2, 3]
		[5, 6, 7, 8]
	*/
	numCopied := copy(s1, s2)
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("num copied: %d\n", numCopied)
	/*
		[5, 6, 7]
		[5, 6, 7, 8]
		This examples shows that the copy function basically will move over as many possible
		records from one slice to another (as much as is allowed be the smaller of the two slices)

		the return value of copy() is the number of records that managed to be copied over.
	*/
	arrExample := [5]int{1, 2, 3, 4, 5}
	// copy only takes slices as arguments, so trying to copy into arrExample won't work
	// to convert arrExample into a slice:

	arrExampleToSlice := arrExample[0:]
	numCopied = copy(arrExampleToSlice, s2)
	fmt.Printf("arrExampleToSlice after copy: %v\n", arrExampleToSlice)
	fmt.Printf("num copied: %d\n", numCopied)
	/*
		arrExampleToSlice after copy: [5, 6, 7, 8, 5]
		num copied: 4

		arrExample got converted to a slice in arrExampleToSlice using the [:] notation
		the copy function copied from s2, into arrExampleToSlice - 4 was the smallest of the two slices
		so 4 records were copied
	*/
}
