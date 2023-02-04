package server

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"GatorLeasing/gator-leasing-server/config"
	"GatorLeasing/gator-leasing-server/handler"
)

type Server struct {
	config       *config.ServerConfig
	leaseHandler *handler.LeaseHandler
}

func NewServer(config *config.ServerConfig, leaseHandler *handler.LeaseHandler) *Server {
	return &Server{
		config:       config,
		leaseHandler: leaseHandler,
	}
}

func (s *Server) Run() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully waits for existing connections to finish -e.g. 15s or 1m")

	httpServer := &http.Server{
		Addr:         s.config.Address,
		Handler:      s.handler(),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	httpServer.Shutdown(ctx)

	log.Println("shutting down")
	os.Exit(0)
}

func (s *Server) handler() *mux.Router {
	r := mux.NewRouter()

	r.Use(mux.CORSMethodMiddleware(r))

	get(r, "/leases", s.leaseHandler.GetAllLeases)
	post(r, "/leases", s.leaseHandler.PostLease)

	r.PathPrefix("/")

	return r
}

// Wrap the router for GET method
func get(router *mux.Router, path string, f func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func post(router *mux.Router, path string, f func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func put(router *mux.Router, path string, f func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func delete(router *mux.Router, path string, f func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc(path, f).Methods("DELETE")
}
