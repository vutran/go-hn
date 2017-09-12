package utils

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func Get(url string) *bytes.Buffer {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var b bytes.Buffer
	if _, err := io.Copy(&b, resp.Body); err != nil {
		log.Fatal(err)
	}

	return &b
}
