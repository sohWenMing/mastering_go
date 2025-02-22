package main

import (
	"flag"
	"fmt"
)

func main() {
	minusK := flag.Bool("k", true, "specify a boolean")
	minusO := flag.Int("o", 1, "specify an integer")

	/* note that flag.Bool and flag.Int (and probably all the return values from of functions that are used to define flags) will
	 return pointer values, so they can't actually be used in the program as their values until
		 flag.Parse is called
		 the pointer values need to also be dereferenced so that the actual values being held by the pointers can be accessed.
	*/

	/*
		flag.Parse() parses the command line flags from os.Args[1:]. This has to be called after all the flags are defined, but before
		the flag values are actually accessed by the program
	*/

	flag.Parse()
	valueK := *minusK
	valueO := *minusO
	fmt.Println("valueK: ", valueK)
	fmt.Println("valueO: ", valueO)
	fmt.Printf("minusK: %v\n", minusK)
	fmt.Printf("minusO %v\n", minusO)
}

/*
	running the following executions

	go run .
	will yield
	valueK: true
	valueO: 1
	minusK: <pointer value>
	minusO: <pointer value>

	if we were to run go run . -h or go run . -help
		Usage of /tmp/go-buiVld2694550581/b001/exe/flags:
		-O int
		specify an integer (default 1)
		-k	specify a boolean (default true)

	the third param of the flag function specifies the help message that would be printed in the event that the program was run with
	the help flag.

	we can also set the flag values when executing the program

	go run . -k false -o=100
*/
