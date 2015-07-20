package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/adamveld12/fortune/quote"
	"github.com/adamveld12/fortune/server"
)

var (
	wait       = flag.Duration("wait", time.Nanosecond, "How long to wait before terminating the process (eg. 1ms, 1.2s)")
	serverType = flag.String("listen", "", "Starts the fortune app as a QOTD server. valid values are http or tcp")
	port       = flag.Int("port", 8080, "")
)

func main() {
	flag.Parse()

	arguments := os.Args
	source := "fortunes.txt"
	if len(arguments) > 1 {
		source = arguments[1]
	}

	if *serverType != "" {
		server.Run(source, *port, server.ServerType(*serverType))
	} else {
		quoteString, err := quote.Find(source)

		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(quoteString)
			time.Sleep(*wait)
		}
	}
}
