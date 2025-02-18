package main

import (
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	programName := filepath.Base(os.Args[0])
	// get the filename based on the first arguement in os.Args, which would default to the program name of this program
	sysLog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, programName)
	/*
		the first param is setting the priority and and also the channel for the syslog logging by using the bit constants
		the secondParam is setting the tag for the new syslog Logger which will the be used in programs
	*/
	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(sysLog)
		// this is setting the output of the main logger to the return value of syslog.New - basically it's saying
		// write to the channel defined as local7 in the syslog configuration
	}
	log.Println("LOG_INFO + LOG_LOCAL7: Logging in Go!")
}
