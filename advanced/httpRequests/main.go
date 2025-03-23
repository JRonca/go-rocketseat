package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	for range 10 {
		resp, err := http.Get("https://www.google.com")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		fmt.Println(resp.Status)
	}
	fmt.Println(time.Since(start))
}
