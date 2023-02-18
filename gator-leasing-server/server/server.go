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
	"github.com/joho/godotenv"

	"GatorLeasing/gator-leasing-server/config"
	"GatorLeasing/gator-leasing-server/handler"
	"GatorLeasing/gator-leasing-server/server/middleware"
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
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

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

	r.Use(middleware.CorsMiddleware)

	handle(r, "/leases", "GET", s.leaseHandler.GetAllLeases, false)
	handle(r, "/leases", "POST", s.leaseHandler.PostLease, true)
	handle(r, "/leases/{id}", "PUT", s.leaseHandler.PutLease, true)
	handle(r, "/leases/{id}", "DELETE", s.leaseHandler.DeleteLease, true)

	r.PathPrefix("/").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			return
		})

	return r
}

func handle(router *mux.Router, path string, method string, f func(w http.ResponseWriter, r *http.Request), requiresAuth bool) {
	if requiresAuth {
		router.Handle(path, middleware.EnsureValidToken()(http.HandlerFunc(f))).Methods(method)
	} else {
		router.Handle(path, http.HandlerFunc(f)).Methods(method)
	}
}
