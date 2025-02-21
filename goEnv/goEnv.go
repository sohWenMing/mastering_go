package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("You are using ", runtime.Compiler, " ")
	fmt.Println("on a", runtime.GOARCH, " machine")
	fmt.Printf("That is running %d CPUs\n", runtime.NumCPU())
	fmt.Printf("And you are currently running %d go routines\n", runtime.NumGoroutine())
}

//version
//CPU nums
//number of goroutines
