package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	rootDir, err := openDir("/")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("val in rootDir %v\n", rootDir)
}

func openDir(path string) (*os.Root, error) {
	err := checkIsDir(path)
	if err != nil {
		return nil, err
	}
	root, err := os.OpenRoot(path)
	if err != nil {
		return nil, err
	}
	return root, nil
}

func checkIsDir(path string) (err error) {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return fmt.Errorf("path specified is not a directory %s", path)
	}
	return nil
}
