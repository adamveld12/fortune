package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func main() {

	fmt.Println(GetQuote())
}

func GetQuote() string {
	rand.Seed(time.Now().UnixNano())
	files, _ := ioutil.ReadDir("./quotes")
	file := files[rand.Intn(len(files))]
	data, _ := ioutil.ReadFile("./quotes/" + file.Name())
	text := strings.Split(string(data), "\n\n")
	line := text[rand.Intn(len(text))]
	return line
}
