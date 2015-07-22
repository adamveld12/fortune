package quote

import "strings"

func Find(source string) (quoteString string, err error) {
	if strings.Contains(source, "http://") || strings.Contains(source, "qotd://") || strings.Contains(source, "tcp://") {
		quoteString, err = tcpService(source)
	} else if hasIndex(source) {
		quoteString, err = fromIndex(source)
	} else {
		quoteString, err = fromFile(source)
	}

	return
}

func GenerateIndex(quoteFile string) error {
	index := createIndexFile(quoteFile)
	if err := writeIndex(quoteFile, index); err != nil {
		return err
	}
	return nil
}
