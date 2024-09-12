package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}
	fmt.Println(resp.Status, resp.Proto)

	// Using custom logWriter to handle the response body
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs)) // Print the response body
	fmt.Println("Just wrote this many bytes: ", len(bs))
	return len(bs), nil
}
