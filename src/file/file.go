package file

import (
	"io"
	"os"
)

// Copy creates a copy for a file
func Copy(source, destination string) error {
	var (
		inputFile  *os.File
		outputFile *os.File
		err        error
	)

	inputFile, err = os.Open(source)
	if err != nil {
		return err
	}

	defer inputFile.Close()

	outputFile, err = os.Create(destination)
	if err != nil {
		return err
	}

	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return err
	}

	return outputFile.Close()
}

// Exists checks if a given filename or directory exists.
func Exists(name string) bool {
	_, err := os.Stat(name)

	return !os.IsNotExist(err)
}
