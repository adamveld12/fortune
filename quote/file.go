package quote

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func fromFile(sourceFilePath string) (string, error) {
	rand.Seed(time.Now().UnixNano())

	text, err := loadFile(sourceFilePath)
	if err != nil {
		return "", err
	}

	line := text[rand.Intn(len(text))]
	return line, nil
}

func loadFile(p string) ([]string, error) {
	data, err := ioutil.ReadFile(p)

	if err != nil {
		return []string{}, errors.New(fmt.Sprint("The file at", p, "does not exist. Specify a file with -f."))
	}

	return strings.Split(string(data), "%"), nil
}
