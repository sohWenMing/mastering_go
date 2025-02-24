package readNumBytes

import (
	"errors"
	"io"
	"os"
)

func readNumBytes(fileName string, numBytes int) (byteSlice []byte, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf := make([]byte, numBytes)
	n, readErr := file.Read(buf)
	// buf is a slice of bytes, which can be passed directly into file.Read because slices are always passed by reference in go
	// n denotes the number of bytes read

	if readErr == io.EOF {
		return nil, err
	}
	if readErr != nil {
		return nil, err
	}
	if n != len(buf) {
		return nil, errors.New("number of bytes read was not equal to what was specifie")
	}
	return buf, nil
}
