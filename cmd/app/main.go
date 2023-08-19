package main

import (
	"net/http"

	"github.com/go-chi/chi"

	handlers "github.com/Pizhlo/wb-L0/internal/app/handlers"
)

func main() {
	//conf := ParseFlags()
	
	run()
}

func run() {
	r := chi.NewRouter()

	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetByID(w, r)
	})
}
