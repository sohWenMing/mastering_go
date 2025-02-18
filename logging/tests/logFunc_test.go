package logFunc_tests

import (
	"bufio"
	"os"
	"strings"
	"testing"

	logFunc "github.com/sohWenMing/mastering_go/logging/log_func"
)

func TestCreateLogger(t *testing.T) {
	path := "./Logfile"
	prefix := "testPrefix "
	file, logger, err := logFunc.CreateLogger(path, prefix)
	if err != nil {
		t.Fatal(err.Error())
	}
	logger.Print("Test log 1")
	logger.Print("Test log 2")

	// if err := file.Sync(); err != nil {
	// 	t.Fatal("Failed to flush file:", err)
	// }
	file.Close()
	expectedStrings := []string{
		"Test log 1",
		"Test log 2",
	}
	i := 0
	file, err = os.Open(path)
	defer file.Close()
	if err != nil {
		t.Fatal("Failed to open file: ", err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineFromFile := scanner.Text()
		if !strings.Contains(lineFromFile, expectedStrings[i]) {
			t.Errorf("line from file %s did not contain expected substr %s",
				lineFromFile, expectedStrings[i])
		}
		if err := scanner.Err(); err != nil {
			t.Fatal(err)
		}
		i++
	}

	if i != len(expectedStrings) {
		t.Errorf("Expected %d lines, file had %d lines\n", len(expectedStrings), i)
	}
	defer func() {
		err := os.Remove(path)
		if err != nil {
			t.Fatal(err)
		}
	}()
}
