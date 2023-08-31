package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Pizhlo/wb-L0/config"
	handler "github.com/Pizhlo/wb-L0/internal/app/handler"
	"github.com/Pizhlo/wb-L0/internal/app/stan/publisher"
	"github.com/Pizhlo/wb-L0/internal/app/stan/subscriber"
	"github.com/Pizhlo/wb-L0/internal/app/storage/cache"
	storage "github.com/Pizhlo/wb-L0/internal/app/storage/postgres"
	"github.com/Pizhlo/wb-L0/internal/service"
	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nats-io/stan.go"
)

func main() {
	// loading config
	conf, err := config.LoadConfig("../..")
	if err != nil {
		log.Fatal("unable to load config: ", err)
	}

	// creating new connection to db
	conn, err := pgxpool.New(context.Background(), conf.DBAddress)
	if err != nil {
		log.Fatal("unable to create connection: ", err)
	}

	// creating db
	db := storage.New(conn)
	defer db.Close()

	// creating cache
	cache := makeCache()

	// creating stan conn, subscriber, publisher
	stanConn, err := stan.Connect("my_cluster", "test")
	if err != nil {
		log.Fatal("unable to connect stan: ", err)
	}

	publisher, err := publisher.New(stanConn)
	if err != nil {
		log.Fatal("unable to create publisher: ", err)
	}

	service := service.New(db, cache)

	log.Println("starting server")

	err = subscriber.New(stanConn, *service)
	if err != nil {
		log.Fatal("unable to create subscriber: ", err)
	}

	err = publisher.SendMsg()
	if err != nil {
		log.Fatal("unable to send msg: ", err)
	}

	handler := handler.NewOrder(*service)

	// starting server
	if err := http.ListenAndServe(":8080", run(service, *handler)); err != nil {
		log.Fatal("error while executing server: ", err)
	}
}

func run(service *service.Service, order handler.Order) chi.Router {
	r := chi.NewRouter()

	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		handler.GetOrderByID(w, r, order)
	})

	return r
}

func makeCache() *cache.Cache {
	delivery := cache.NewDelivery()
	payment := cache.NewPayment()
	item := cache.NewItem()
	order := cache.NewOrder()

	return &cache.Cache{
		Order:    order,
		Delivery: delivery,
		Item:     item,
		Payment:  payment,
	}
}
