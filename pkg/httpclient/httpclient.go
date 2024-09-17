package httpclient

import (
	"io"
	"log"
	"net/http"
)

func Request(method string, url string) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Fatal(err)
	}

	body, readError := io.ReadAll(req.Body)

	if readError != nil {
		log.Fatal(readError)
	}

	return body, nil
}
