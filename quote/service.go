package quote

import (
	"errors"
	"log"
)

func TCPService(url string) (string, error) {
	log.Println("TCP quote service not implemented.")
	return "", errors.New("TCP quote service not implemented.")
}

func HttpService(url string) (string, error) {
	log.Println("HTTP quote service not implemented.")
	return "", errors.New("HTTP quote service not implemented.")
}
