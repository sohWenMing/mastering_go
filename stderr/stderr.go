package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	myString := ""
	args := os.Args
	if len(args) == 1 {
		myString = "Please give me one more argument!"
	}
	io.WriteString(os.Stdout, "This message is comign from standard output\n")
	io.WriteString(os.Stderr, fmt.Sprintf("%s\n", myString))
}

/*
	if we were to run this program, it would be impossible to figure out whether the message was coming from stdout or stderr
	however, there is a way we can write the information from either stdout or stderr to a file

	go run . > /tmp/stdout would write everything that would be coming from stdout into a temp file called stdout
	// go run . 2> /tmp/stdError would write everything that was coming from stdErr into a temp file called stdError

	*[main][~/workspace/github.com/sohWenMing/mastering_go/stderr]$ go run . > /tmp/stdout 2> /tmp/stdError
	*[main][~/workspace/github.com/sohWenMing/mastering_go/stderr]$ cat /tmp/stdError
	Please give me one more argument!
	*[main][~/workspace/github.com/sohWenMing/mastering_go/stderr]$ cat /tmp/stdout
	This message is comign from standard output

*/
