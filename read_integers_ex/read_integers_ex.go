package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Enter integers to have them printed back to you!")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "STOP" {
			fmt.Println("Stop was entered, exiting program")
			os.Exit(0)
		}
		integer, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("input was not an integer")
			continue
		}
		fmt.Println("integer entered: ", integer)
	}
}
