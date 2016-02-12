package utilities

import (
	"bufio"
	"log"
	"os"
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
