package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please give me one more floats!")
		os.Exit(1)
	}
	min, err := parseArgToFloat(args[1])
	if err != nil {
		log.Fatal(err)
	}
	max, err := parseArgToFloat(args[1])
	if err != nil {
		log.Fatal(err)
	}
	for _, val := range args[1:] {
		floatVal, err := parseArgToFloat(val)
		if err != nil {
			log.Fatal(err)
		}
		if floatVal < min {
			min = floatVal
		}
		if floatVal > max {
			max = floatVal
		}
	}
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}

func parseArgToFloat(arg string) (float64, error) {
	val, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		return 0.0, err
	}
	return val, nil
}
