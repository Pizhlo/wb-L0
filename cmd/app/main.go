package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// loading config
	conf, err := config.LoadConfig("../..")
	if err != nil {
		log.Fatal("unable to load config: ", err)
	}

	// creating new connection to db
	conn, err := pgxpool.New(serverCtx, conf.DBAddress)
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

	publisher := publisher.New(stanConn)

	service := service.New(db, cache)

	log.Println("starting server")

	err = subscriber.New(stanConn, *service)
	if err != nil {
		log.Fatal("unable to create subscriber: ", err)
	}

	ticker := startTicker(serverCtx, conf.Ticker)
	done := make(chan bool)

	err = publisher.Start(ticker, done)
	if err != nil {
		log.Fatal("unable to send msg: ", err)
	}

	handler := handler.NewOrder(*service)

	err = service.Recover(serverCtx)
	if err != nil {
		log.Fatal("unable to recover data: ", err)
	}

	server := &http.Server{Addr: "0.0.0.0:8080", Handler: router(service, *handler)}

	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)
		defer cancel()

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal("error while shutdown server: ", err)
		}
		serverStopCtx()
	}()

	// starting server
	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal("error while starting server: ", err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()

}

func router(service *service.Service, order handler.Order) http.Handler {
	r := chi.NewRouter()

	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		handler.GetOrderByID(w, r, order)
	})

	return r
}

func makeCache() *cache.Cache {
	order := cache.NewOrder()

	return cache.New(order)
}

func startTicker(ctx context.Context, duration time.Duration) *time.Ticker {
	fmt.Println("starting ticker with duration", duration)
	ticker := time.NewTicker(duration)

	return ticker
}
