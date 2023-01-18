package main

import (
	"fmt"
	"io"
	"net/http"
)

var pr = fmt.Printf

func main() {
	ctype, err := contentType("https://catfact.ninja/fact")
	if (err != nil) {
		pr("ERROR: %v \n", err)
	}

	pr("content type is %v \n", ctype)
}

func contentType(url string) (string, error) {
	response, err := http.Get(url)
	if (err != nil) {
		pr("ERROR: %v \n", err)
		return "", err
	}
	defer response.Body.Close()

	headers := response.Header.Get("Content-Type")
	body, err := io.ReadAll(response.Body)
	if (err != nil) {
		pr("ERROR: %v \n", err)
	}
	pr("response %v \n",)

	return headers, nil
}
