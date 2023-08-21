package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"

	handlers "github.com/Pizhlo/wb-L0/internal/app/handlers"
	storage "github.com/Pizhlo/wb-L0/internal/app/storage/postgres"
	stream "github.com/Pizhlo/wb-L0/internal/stream"
	"github.com/Pizhlo/wb-L0/internal/stream/publisher"
	"github.com/Pizhlo/wb-L0/internal/stream/subscriber"
	"github.com/Pizhlo/wb-L0/service"
)

func main() {
	conf := ParseFlags()

	conn, err := storage.Connect(conf.DBAddress)
	if err != nil {
		log.Fatal("unable to connect db: ", err)
	}

	db := storage.New(conn, conf.DBAddress, conf.MigratePath)
	// if err != nil {
	// 	log.Fatal("unable to create db: ", err)
	// }

	if db != nil {
		defer db.Close()
	}

	nc, err := stream.Connect()
	if err != nil {
		log.Fatal("unable to connect nats: ", err)
	}

	if nc != nil {
		defer nc.Close()
	}

	// err = nc.Publish("foo", []byte("check conntection")); if err != nil {
	// 	log.Fatal("unable to send msg: ", err)
	// }

	publisher, err := publisher.New(nc.Conn)
	if err != nil {
		log.Fatal("unable to create publisher: ", err)
	}
	if publisher != nil {
		defer publisher.EncodedConn.Close()
	}

	subscriber := subscriber.New(publisher.EncodedConn, db)
	// if err != nil {
	// 	log.Fatal("unable to create subscriber: ", err)
	// }

	err = publisher.SendMsg()
	if err != nil {
		log.Fatal("unable to send msg: ", err)
	}
	
	service := service.New(db, publisher, subscriber)

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
