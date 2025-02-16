package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	args := os.Args
	/*
		args will get the arguments that were run from the command line at the point that a program is executed
		by default, the name of the program that is run is the first argument, args[0]
	*/
	if len(args) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = args[1]
	}

	io.WriteString(os.Stdout, myString)
	/*
		io.WriteString takes in the the type io.Writer as the first argument - os.Stdout fulfils this interface because it ahs the
		Write() method attached to it.


		running go run stdout.go will print "Please give me one argument!"
		running go run stdout.go 123 will print "123"
	*/
	io.WriteString(os.Stdout, "\n")
}
