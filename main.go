package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

var serverUrl string
var wait time.Duration
var filePath string

func init() {
	flag.StringVar(&filePath, "file", "fortunes.txt", "Retrieves a quote from the specified quote file.")
	flag.StringVar(&serverUrl, "server", "", "Retrieves a quote from the specified quote server.")
	flag.DurationVar(&wait, "wait", time.Nanosecond, "How long to wait before terminating the process (eg. 1ms, 1.2s).")
}

func main() {
	flag.Parse()

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
	return ""
}

func GetQuoteFromFile() string {

	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err.Error())
	}

	text := strings.Split(string(data), "%")

	rand.Seed(time.Now().UnixNano())
	line := text[rand.Intn(len(text))]

	return line
}
