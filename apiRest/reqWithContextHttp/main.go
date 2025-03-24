package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://reqres.in/api/users", nil)

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
