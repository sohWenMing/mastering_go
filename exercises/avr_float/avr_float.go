package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("no arguments were entered, exiting program...")
		os.Exit(1)
	}
	args := os.Args[1:]
	floatArgs := []float64{}
	for _, arg := range args {
		floatArg, err := strconv.ParseFloat(arg, 64)
		if err == nil {
			floatArgs = append(floatArgs, floatArg)
		}
	}
	if len(floatArgs) == 0 {
		fmt.Println("no floats were entered, exiting program...")
		os.Exit(1)
	}
	var sum float64
	for _, floatArg := range floatArgs {
		sum += floatArg
	}
	fmt.Printf("average of all floats entered: %.2f\n", sum/float64(len(floatArgs)))
}
