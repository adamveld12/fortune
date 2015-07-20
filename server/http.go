package server

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/adamveld12/fortune/quote"
)

func fromHttp(port int, quoteFile string) error {
	http.HandleFunc("/", handler(quoteFile))
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func handler(quoteFilePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		quoteString, err := quote.Find(quoteFilePath)
		if err != nil {
			io.WriteString(w, "Sorry, a quote could not be retrieved. Please try again later.")
			log.Println(err)
		} else {
			io.WriteString(w, quoteString)
		}
	}
}
