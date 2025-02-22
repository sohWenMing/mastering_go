package main

import (
	"flag"
	"fmt"
	"os"

	fileReading "github.com/sohWenMing/mastering_go/io/reading"
)

func main() {

	action := flag.String("action", "", "action: read|write")
	fileName := flag.String("filename", "", "enter path of filename")
	flag.Parse()

	enteredAction := *action
	if enteredAction == "" {
		fmt.Println("flag for action is required, usage: readOrWrite -action=<read|write>")
		os.Exit(1)
	}
	// user entering either "read" or "write" is required, if none should early exit
	switch enteredAction {
	case "read":
		err := fileReading.ReadFileByLine(*fileName)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	default:
		fmt.Println("action entered does is not supported - run readOrWrite -help for usage")
		os.Exit(1)
	}
}
