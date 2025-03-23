package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	const n = 10
	var wg sync.WaitGroup
	wg.Add(n)

	// criando contexto
	ctx := context.Background()
	// criando um timeout de 5 segundos para o contexto
	// Essa função retorna um novo contexto com o timeout e uma função de cancelamento
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	// defer cancel para garantir que o cancelamento seja chamado
	defer cancel()

	// Criando servidor http para gerar o timeout
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Second)
		fmt.Println("Hello World!")
	}))

	for range n {
		// passando o contexto como argumento
		go func(ctx context.Context) {
			defer wg.Done()

			// Vamos criar um request
			req, err := http.NewRequestWithContext(
				ctx,
				"GET",
				server.URL,
				nil,
			)

			if err != nil {
				panic(err)
			}

			// Passando o request para ser executado pelo client
			resp, err := http.DefaultClient.Do(req)

			if err != nil {
				// Verificando se o erro é um erro de contexto de timeout
				if errors.Is(err, context.DeadlineExceeded) {
					fmt.Println("Timeout")
					return
				}

				panic(err)
			}

			defer resp.Body.Close()

		}(ctx)
	}

	wg.Wait()
	fmt.Println(time.Since(start))
}
