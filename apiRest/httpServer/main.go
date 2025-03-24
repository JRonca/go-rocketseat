package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

func main() {

	// criando um mux para rotear as requisições
	mux := http.NewServeMux()
	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world")
	})

	// criando um servidor http com algumas configurações de porta e timeout
	srv := &http.Server{
		Addr:                         ":8080",
		Handler:                      mux,
		DisableGeneralOptionsHandler: false,
		ReadTimeout:                  10 * time.Second,
		WriteTimeout:                 10 * time.Second,
		IdleTimeout:                  1 * time.Minute,
	}

	// iniciando o servidor
	// validando erros, exceto erro de servidor fechado
	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}

}
