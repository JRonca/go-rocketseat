package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()
		next.ServeHTTP(w, r)
		fmt.Println(r.URL.String(), r.Method, time.Since(begin))
	})
}

func main() {

	// criando um mux para rotear as requisições com o método GET
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world")
	})

	mux.HandleFunc("GET /api/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "Hello user %s", id)
	})

	// criando um servidor http com algumas configurações de porta e timeout
	srv := &http.Server{
		Addr:                         ":8080",
		Handler:                      Log(mux),
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
