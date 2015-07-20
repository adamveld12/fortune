package quote

import "strings"

func Find(source string) (quoteString string, err error) {
	if strings.Contains(source, "http://") || strings.Contains(source, "qotd://") || strings.Contains(source, "tcp://") {
		quoteString, err = tcpService(source)
	} else {
		quoteString, err = file(source)
	}
	return
}
