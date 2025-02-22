package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var manyNames NamesFlag
	minusK := flag.Int("k", 0, "An int")
	minusO := flag.String("o", "Wen", "The name")
	flag.Var(&manyNames, "names", "Comma-separated list")

	/*
		Here flag.Var is taking nameNames in as the Value interface param int the flag.Var function

		we're basically setting flag called names which we can enter a comma separated list into, and then whatever is entered
		as that comma separated list will then be stored in the variable that is defined in the value param. see below for basic
		reference of the Var function, from the flag package
	*/

	flag.Parse()
	fmt.Println("-k:", *minusK)
	fmt.Println("-o:", *minusO)

	for i, item := range manyNames.GetNames() {
		fmt.Println(i, item)
	}

	fmt.Println("Remaning command-line arguments:")
	for index, val := range flag.Args() {
		fmt.Println(index, ":", val)
	}
}

type NamesFlag struct {
	Names []string
}

/*
	NamesFlag is an instance of the Value interface, which is required to be used in the flag.Var() function

	## definition of the Var function ##
	func Var(value Value, name string, usage string)

		type value interface {
		String() string
		Set(string) error
	}
*/

func (s *NamesFlag) GetNames() []string {
	return s.Names
}

// returns the slice of strings in the Names field in NamesFlag

func (s *NamesFlag) String() string {
	return fmt.Sprint(s.Names)
}

//prints a string representation of the slice of strings in the Names field in NamesFlag

func (s *NamesFlag) Set(v string) error {
	if len(s.Names) > 0 {
		return fmt.Errorf("cannot use names flag more than once")
	}
	//the Names field should never bet set more than once, if there's already at least one entry in slice in field, return error
	names := strings.Split(v, ",")
	s.Names = append(s.Names, names...)
	return nil
}
