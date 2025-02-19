package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("no arguments were entered into program - exiting...")
		os.Exit(1)
	}
	argsTocalc := os.Args[1:]
	intArgs := []int{}
	floatArgs := []float64{}

	for _, arg := range argsTocalc {
		integer, err := strconv.Atoi(arg)
		if err == nil {
			intArgs = append(intArgs, integer)
			continue
		}
		float, err := strconv.ParseFloat(arg, 64)
		if err == nil {
			floatArgs = append(floatArgs, float)
		}
	}
	var sum float64
	for _, val := range intArgs {
		sum += float64(val)
	}
	for _, val := range floatArgs {
		sum += val
	}
	numArgs := len(intArgs) + len(floatArgs)
	fmt.Printf("average of all arguments to 2 decimal places: %.2f", sum/float64(numArgs))
}
