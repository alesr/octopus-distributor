package utilities

import (
	"bufio"
	"errors"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
	// ErrRandomFuncError used at Random()
	ErrRandomFuncError = errors.New("Random: function expect an int arg greater than zero")
)

// LoadFile given a file path returns the file content
func LoadFile(filepath string) ([]string, error) {

	// checks if filepath is valid
	if _, err := checkFile(filepath); err != nil {
		return nil, err
	}

	// open the file
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	// read file line by line and append each line to slice content
	content := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	// check for errors
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return content, nil
}

func checkFile(filepath string) (bool, error) {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

// Random returns a pseudo random number between zero and max
func Random(max int) (int, error) {
	// it means that the source slice is empty.
	if max <= 0 {
		return 0, ErrRandomFuncError
	}

	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max), nil
}
