package logFunc

import (
	"log"
	"os"
)

func CreateLogger(path, prefix string) (file *os.File, logger *log.Logger, err error) {
	file, err = os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return nil, nil, err
	}

	logger = log.New(file, prefix, log.Ldate|log.Ltime|log.LUTC|log.Lshortfile)
	return file, logger, nil
}
