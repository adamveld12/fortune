package server

import (
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/adamveld12/fortune/quote"
)

type Listener func(int, string) error

func fromTcp(port int, quoteFile string) error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatal(err)
		return errors.New(fmt.Sprintf("Could not listen on port %d.", port))
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Error accepting: ", err.Error())
		} else {
			go handleConn(quoteFile, conn)
		}
	}
}

func handleConn(quoteFile string, conn net.Conn) {
	defer conn.Close()
	quote, err := quote.Find(quoteFile)

	if err != nil {
		log.Println("Error occurred reading quote from file: ", quoteFile)
		conn.Write([]byte("Sorry, a quote could not be retrieved. Please try again later."))
	} else {
		conn.Write([]byte(quote))
	}

}
