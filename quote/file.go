package quote

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// File loads a random quote from the file p and returns it
// checks if an index file exists as "{p}.index", and if so uses it
// Returns an error if the file cannot be opened, or if the format is not correct
func File(p string) (string, error) {
	// check if an index file exists
	return readFromRawFile(p)
}

// GeneratesIndex Generates  an index file for the file p and saves it at the app's current directory
// Returns an error if the index file could not be written to, or if the source file is not correct
func GenerateIndex(p string) (string, error) {
	return "", nil
}

func hasIndex(p string) bool {
	return false
}

func readFromIndex(p string) (string, error) {
	return "", nil
}

func readFromRawFile(p string) (string, error) {
	data, err := ioutil.ReadFile(p)

	if err != nil {
		return "", errors.New(fmt.Sprint("The file at", p, "does not exist. Specify a file with -f."))
	}

	text := strings.Split(string(data), "%")

	rand.Seed(time.Now().UnixNano())
	line := text[rand.Intn(len(text))]

	return line, nil
}
