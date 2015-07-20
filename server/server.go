package server

import (
	"log"
	"os"
	"strconv"
)

type ServerType string

const (
	HTTPserver ServerType = "http"
	TCPserver  ServerType = "tcp"
)

func Run(source string, port int, serverType ServerType) {
	defer func() { log.Println("Server shutting down...") }()

	if envPort := os.Getenv("PORT"); envPort != "" && port > 0 {
		parsedPort, err := strconv.Atoi(envPort)
		port = parsedPort
		if err != nil {
			log.Fatal("A number must be passed to -port.")
		}
	} else {
		port = 8080
	}

	log.Printf("Being philosophical over %s on port %d.", serverType, port)

	var err error
	if serverType == TCPserver {
		err = fromTcp(port, source)
	} else if serverType == HTTPserver {
		err = fromHttp(port, source)
	}

	if err != nil {
		log.Fatal(err)
	}

}
