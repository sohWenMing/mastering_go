package main

import (
	"fmt"

	"github.com/mactsouk/go/simpleGitHub"
)

/*
	if this program is run without first importing the required files, the following will print in the console

	downloading_packages.go:6:2: no required module provides package github.com/mactsouk/go/simpleGitHub; to add it:
        go get github.com/mactsouk/go/simpleGitHub

	the required dependencies are downloaded into the module cache, which in general will be located in ~/go/pkg/mod

	these files are the available to all later go projects, so there is no need to download the required dependencies on a per project basis

	In the event of downloading someone else's project, running go mod tidy will download all required dependencies into the local machine
	for further usage

*/

func main() {
	fmt.Println(simpleGitHub.AddTwo(5, 6))
}
