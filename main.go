package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/adamveld12/fortune/quote"
	"github.com/adamveld12/fortune/server"
)

var (
	generateIndex = flag.Bool("generateIndex", false, "Generates an index of the specified quote file")
	serverType    = flag.String("listen", "", "Starts the fortune app as a QOTD server. valid values are http or tcp")
	wait          = flag.Duration("wait", time.Nanosecond, "How long to wait before terminating the process (eg. 1ms, 1.2s)")
	port          = flag.Int("port", 8080, "")
)

func main() {
	flag.Parse()

	arguments := flag.Args()
	source := "fortunes.txt"

	if len(arguments) > 1 {
		source = arguments[1]
	}

	if *serverType != "" {
		server.Run(source, *port, server.ServerType(*serverType))
	} else if *generateIndex {
		fmt.Printf("Generated index %s.index\n", source)
		if err := quote.GenerateIndex(source); err != nil {
			log.Fatal(err)
		}
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
