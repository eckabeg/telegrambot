package main

import (
	"bytes"
	"net/http"
)

func download(url string) []byte {
	response, errs := http.Get(programmDownloadUrl + url)
	if errs != nil {
		print(errs)
	}
	buf := bytes.NewBuffer(make([]byte, 0, response.ContentLength))
	_, err := buf.ReadFrom(response.Body)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}
