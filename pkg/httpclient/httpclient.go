package httpclient

import (
	"io"
	"log"
	"net/http"
)

func Request(method string, url string) ([]byte, error) {
	req, _ := http.NewRequest(method, url, nil)

	response, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	body, readError := io.ReadAll(response.Body)

	if readError != nil {
		log.Fatal(readError)
	}

	return body, nil
}
