package quote

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func createIndexFile(quotePath string) *Index {
	f, err := os.Open(quotePath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	qr := bufio.NewReader(f)

	position, entries := int64(0), []IndexEntry{}

	for {
		line, err := qr.ReadBytes('%')

		if err != nil && err != io.EOF && len(line) > 0 {
			log.Println("error reading entry from quote file")
			log.Fatal(err)
		}

		quoteEntryByteSize := int64(binary.Size(line))
		entries = append(entries, IndexEntry{Location: position, Size: quoteEntryByteSize - 1})
		position += quoteEntryByteSize

		if err == io.EOF {
			break
		}
	}

	return &Index{Count: int32(len(entries)), Entries: entries}
}

func writeIndex(quoteFile string, index *Index) error {
	out, err := os.Create(indexFileName(quoteFile))

	if err != nil {
		return err
	}

	defer out.Close()

	encoder := gob.NewEncoder(out)

	if err = encoder.Encode(index); err != nil {
		return err
	}

	return nil
}

func fromIndex(p string) (string, error) {
	index := openIndex(p)

	//generate a rand number between 0 -> length
	rand.Seed(time.Now().UnixNano())
	quoteIdx := rand.Int31n(index.Count)

	f, err := os.Open(p)
	if err != nil {
		return "", err
	}

	entry := index.Entries[quoteIdx]
	offset, err := f.Seek(entry.Location, 0)
	if err != nil {
		return "", err
	}

	buf := make([]byte, entry.Size)
	if _, err := f.ReadAt(buf, offset); err != nil {
		return "", err
	}

	return string(buf), nil
}

func indexFileName(quoteFile string) string {
	dir, name := filepath.Split(quoteFile)
	indexFileName := filepath.Join(dir, name+".index")
	return indexFileName
}

func openIndex(p string) *Index {
	indexFileName := indexFileName(p)
	data, err := ioutil.ReadFile(indexFileName)

	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.NewBuffer(data)

	decoder := gob.NewDecoder(buf)

	var idxStruct Index

	if err = decoder.Decode(&idxStruct); err != nil {
		log.Fatal(err)
	}

	return &idxStruct
}

func hasIndex(p string) bool {
	fname := indexFileName(p)
	f, err := os.Open(fname)

	if err == nil {
		defer f.Close()
	}

	return err == nil
}

type Index struct {
	Count   int32
	Entries []IndexEntry
}

type IndexEntry struct {
	Location int64
	Size     int64
}
