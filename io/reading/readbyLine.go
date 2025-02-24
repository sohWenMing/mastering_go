package fileReading

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadFileByLine(fileName string) (err error) {
	fmt.Printf("filename entered: %s\n", fileName)
	file, err := openFile(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	//always defer closing of the file, unless the file still needs to be accessed outside of the function

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error read file %s, %v", fileName, err)
		}
		fmt.Println(line)
	}
	return nil
}

func openFile(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}
