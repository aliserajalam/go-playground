package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Get request
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	stringBody := string(bytes)
	fmt.Println(stringBody)
	resp.Body.Close()
}
