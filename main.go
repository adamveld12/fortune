package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var serverUrl, filePath, serverType string
var port int
var wait time.Duration

func init() {
	flag.StringVar(&filePath, "file", "fortunes.txt", "Retrieves a quote from the specified quote file.")
	flag.StringVar(&serverUrl, "server", "", "Retrieves a quote from the specified quote server.")
	flag.DurationVar(&wait, "wait", time.Nanosecond, "How long to wait before terminating the process (eg. 1ms, 1.2s).")
	flag.StringVar(&serverType, "listen", "", "Starts the fortune app as a QOTD server. valid values are http or tcp")
	flag.IntVar(&port, "port", 8080, "")
}

func main() {
	flag.Parse()

	if envPort := os.Getenv("PORT"); port == 8080 && envPort != "" {
		port, err := strconv.Atoi(envPort)
		if err != nil {
			log.Fatal("A number must be passed to -port.")
		}
	}

	if serverType == "tcp" {
		runTcpServer(port)
	} else if serverType == "http" {
		runHttpServer(port)
	} else if serverType != "" {
		log.Fatal("listen argument is incorrect. Valid values are \"tcp\" and \"http\"")
	}

	var quote string

	if serverUrl != "" {
		quote = GetQuoteFromService()
	} else {
		quote = GetQuoteFromFile()
	}

	fmt.Println(quote)
	time.Sleep(wait)
}

func GetQuoteFromService() string {
	// dial via tcp to the QOTD service
	// the spec says that we should receive
	// a quote immediately on connect
	return ""
}

func GetQuoteFromFile() string {

	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Println("The specified file at", filePath, "does not exist.")
		fmt.Println("Specify a file with -f.")
		return ""
	}

	text := strings.Split(string(data), "%")

	rand.Seed(time.Now().UnixNano())
	line := text[rand.Intn(len(text))]

	return line
}
