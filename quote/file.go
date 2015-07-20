package quote

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// File loads a random quote from the file p and returns it
// checks if an index file exists as "{p}.index", and if so uses it
// Returns an error if the file cannot be opened, or if the format is not correct
func file(p string) (string, error) {
	// check if an index file exists
	return readFromRawFile(p)
}

func hasIndex(p string) bool {
	dir, name := filepath.Split(p)
	indexFileName := filepath.Join(dir, strings.TrimSuffix(name, filepath.Ext(name))+".index")
	f, err := os.Open(indexFileName)

	if err == nil {
		defer f.Close()
	}

	return os.IsExist(err)
}

// GeneratesIndex Generates an index file for the file p and saves it at the app's current directory
// Returns an error if the index file could not be written to, or if the source file is not correct
func GenerateIndex(p string) error {
	file, err := os.Open(p)

	if err != nil {
		return errors.New(fmt.Sprint("The file at", p, "does not exist. Specify a file with -f."))
	}

	defer file.Close()

	dir, name := filepath.Split(p)
	indexFileName := filepath.Join(dir, strings.TrimSuffix(name, filepath.Ext(name))+".index")

	indexFile, err := os.Create(indexFileName)

	if err != nil {
		return errors.New(fmt.Sprint("The file at", p, "does not exist. Specify a file with -f."))
	}

	log.Println("writing index file to ", indexFileName)

	defer indexFile.Close()

	reader := bufio.NewReader(file)
	position := 0
	count := 0

	for {
		line, err := reader.ReadBytes('%')

		if err != nil {
			return err
		}

		binary.Write(indexFile, binary.LittleEndian, position)
		binary.Write(indexFile, binary.LittleEndian, len(line))
		position += len(line)
		count++
	}

	binary.Write(indexFile, binary.LittleEndian, count)

	return nil
}

func readFromIndex(p string) (string, error) {
	indexFile := openIndex(p)
	defer indexFile.Close()

	// read length
	//	quoteCount := indexQuoteCount(indexFile)

	// generate a rand number between 0 -> length
	//	rand.Seed(time.Now().UnixNano())
	//	quoteIdx := rand.Int32n()
	// read index and length in the index from that position
	// open actual quote file
	// read quote starting at pos

	return "", nil
}

func indexQuoteCount(indexFile *os.File) int {
	size := binary.Size(4)
	fileLengthIdx, err := indexFile.Seek(int64(size), 2)
	if err != nil {
		log.Fatal(err)
	}

	fileLengthBytes := make([]byte, size)
	_, err = indexFile.ReadAt(fileLengthBytes, fileLengthIdx)
	if err != nil {
		log.Fatal(err)
	}

	var fileLength int
	err = binary.Read(indexFile, binary.LittleEndian, fileLength)
	if err != nil {
		log.Fatal(err)
	}
	return fileLength
}

func openIndex(p string) *os.File {
	dir, name := filepath.Split(p)
	indexFileName := filepath.Join(dir, strings.TrimSuffix(name, filepath.Ext(name))+".index")
	indexFile, err := os.Open(indexFileName)

	if err != nil {
		log.Fatal(err)
	}
	return indexFile
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
