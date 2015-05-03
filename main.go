package main

import (
	"flag"
	"fmt"
	"github.com/adamveld12/fortune/quote"
	"github.com/adamveld12/fortune/server"
	"log"
	"os"
	"strconv"
	"time"
)

var serverUrl, filePath, serverType string
var port int
var wait time.Duration

func init() {
	flag.StringVar(&filePath, "file", "fortunes.txt", "Retrieves a quote from the specified quote file.")
	flag.StringVar(&serverUrl, "server", "", "Retrieves a quote from the specified quote server.")
	flag.DurationVar(&wait, "wait", time.Nanosecond, "How long to wait before terminating the process (eg. 1ms, 1.2s).")
	flag.StringVar(&serverType, "listen", "nil", "Starts the fortune app as a QOTD server. valid values are http or tcp")
	flag.IntVar(&port, "port", 8080, "")
}

func main() {
	flag.Parse()

	if serverType != "nil" {

		if envPort := os.Getenv("PORT"); port == 8080 && envPort != "" {
			parsedPort, err := strconv.Atoi(envPort)
			port = parsedPort
			if err != nil {
				log.Fatal("A number must be passed to -port.")
			}
		}

		log.Printf("%s on port %d started.", serverType, port)

		var err error
		if serverType == "tcp" {
			err = server.Tcp(port, filePath)
		} else if serverType == "http" {
			err = server.Http(port, filePath)
		}

		if err != nil {
			log.Fatal(err)
		}

		log.Println("Server shutting down...")
	} else {
		var quoteString string
		var err error

		if serverUrl != "" {
			quoteString = GetQuoteFromService()
		} else {
			quoteString, err = quote.File(filePath)
		}

		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(quoteString)
			time.Sleep(wait)
		}
	}
}

func GetQuoteFromService() string {
	// dial via tcp to the QOTD service
	// the spec says that we should receive
	// a quote immediately on connect
	return ""
}
