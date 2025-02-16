package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
	/*
		the scanner is reading from os.Stdin, which itself is a pointer to an os.File

		func NewScanner takes in the io.Reader interface as an arg

		to fulfil the io.Reader interface, the type needs to have a method attaced to it of:
		Read(p []byte) (n int, err error) - n will be the number of bytes read and err will be returned if there is a error


		the scanner will keep reading lines from stdin until it gets an EOF signal. This can be signified by hitting CTRL+D in the terminal
		which sends the signal to the program.
	*/
}
