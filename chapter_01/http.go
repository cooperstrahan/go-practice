package main

import (
	"fmt"
	"io"
	"net/http"
)

func connect_http() {
	// simple http request to example.com
	resp, _ := http.Get("http://example.com/")
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	resp.Body.Close()
}
