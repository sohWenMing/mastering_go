package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("no arguments passed into which program, exiting program")
		os.Exit(0)
	}
	if len(args) > 2 {
		fmt.Println("which program only accepts one filename, program terminating")
		os.Exit(1)
	}
	searchString := os.Args[1]
	fmt.Printf("search string entered: %s\n", searchString)
	foundPaths, err := openPathForReading("/", searchString)
	if err != nil {
		fmt.Println("error returned from program", err.Error())
		os.Exit(1)
	}
	for i, path := range foundPaths {
		fmt.Printf("found result %d: %s\n", i, path)
	}
	os.Exit(0)
}

func openPathForReading(path string, searchString string) (foundPaths []string, err error) {
	foundPaths = []string{}
	//init foundPaths to an empty slice before doing anything else

	info, err := os.Stat(path)
	if err != nil {
		{
			return foundPaths, err
		}
		// if os.Stat returns an error, then there should be no foundPaths to return, just return empty slice and the assigned error
	}
	name := info.Name()
	switch info.IsDir() {

	case false:
		evalPath := fmt.Sprintf("%s/%s", path, name)
		if name == searchString {
			foundPaths = append(foundPaths, evalPath)
		}
		//only if the name of the file matches the searchString, then append it to foundPaths

	case true:
		entriesInDir, err := os.ReadDir(path)
		if err != nil {
			return foundPaths, err
		}
		for _, entryInDir := range entriesInDir {
			evalPath := fmt.Sprintf("%s/%s", path, entryInDir.Name())
			if !entryInDir.IsDir() {
				if entryInDir.Name() == searchString {
					foundPaths = append(foundPaths, evalPath)
				}
				continue
			} else {
				foundPathsFromSearch, err := openPathForReading(evalPath, searchString)
				if err != nil {
					continue
				}
				foundPaths = append(foundPaths, foundPathsFromSearch...)

			}
		}
	}
	return foundPaths, nil
}
