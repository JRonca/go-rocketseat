package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	postReqToReqRes()
	getReqToReqRes()
}

func postReqToReqRes() {
	resp, err := http.Post("https://reqres.in/api/users", "application/json", nil)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	fmt.Println(string(data))
}

func getReqToReqRes() {
	resp, err := http.Get("https://reqres.in/api/users")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	fmt.Println(string(data))
}
