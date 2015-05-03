package server

import (
	"fmt"
	"github.com/adamveld12/fortune/quote"
	"io"
	"log"
	"net/http"
)

func Http(port int, quoteFile string) error {
	http.HandleFunc("/", handler(quoteFile))
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func handler(quoteFilePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		quoteString, err := quote.File(quoteFilePath)
		if err != nil {
			io.WriteString(w, "Sorry, a quote could not be retrieved. Please try again later.")
			log.Println(err)
		} else {
			io.WriteString(w, quoteString)
		}
	}
}
