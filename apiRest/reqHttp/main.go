package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "https://reqres.in/api/users", nil)

	if err != nil {
		panic(err)
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("authorization", "123456")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
