package quote

import (
	"errors"
	"log"
)

func tcpService(url string) (string, error) {
	log.Println("TCP quote service not implemented.")
	return "", errors.New("TCP quote service not implemented.")
}

func httpService(url string) (string, error) {
	log.Println("HTTP quote service not implemented.")
	return "", errors.New("HTTP quote service not implemented.")
}
