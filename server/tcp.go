package server

import (
	"errors"
	"github.com/adamveld12/fortune/quote"
	"log"
	"net"
)

type Listener func(int, string) error

func Tcp(port int, quoteFile string) error {
	l, err := net.Listen("tcp", string(port))

	if err != nil {
		return errors.New("Could not listen on port " + string(port) + ".")
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Error accepting:", err.Error())
		}

		go func() {
			defer conn.Close()
			quote, err := quote.File(quoteFile)

			if err != nil {
				log.Println("Error occurred reading quote from file", quoteFile)
			}

			conn.Write([]byte(quote))
		}()
	}
}
