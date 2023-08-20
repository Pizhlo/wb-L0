package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"

	handlers "github.com/Pizhlo/wb-L0/internal/app/handlers"
	storage "github.com/Pizhlo/wb-L0/internal/app/storage/postgres"
	"github.com/Pizhlo/wb-L0/service"
)

func main() {
	conf := ParseFlags()

	conn, err := storage.Connect(conf.DBAddress)
	if err != nil {
		log.Fatal("unable to connect db: ", err)
	}

	db := storage.New(conn, conf.DBAddress,conf.MigratePath)
	// if err != nil {
	// 	log.Fatal("unable to create db: ", err)
	// }

	if db != nil {
		defer db.Close()
	}

	service := service.New(db)

	if err := http.ListenAndServe(":8080", run(service)); err != nil {
		log.Fatal("error while executing server: ", err)
	}
}

func run(service *service.Service) chi.Router {
	r := chi.NewRouter()

	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetOrderByID(w, r, *service)
	})

	return r
}
