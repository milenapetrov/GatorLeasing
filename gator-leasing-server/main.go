package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully waits for existing connections to finish -e.g. 15s or 1m")

	r := mux.NewRouter()

	r.HandleFunc("/leases", getLeases).Methods(http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions)

	r.Use(mux.CORSMethodMiddleware(r))

	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)

	log.Println("shutting down")
	os.Exit(0)
}

type lease struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	BuildingType string `json:"buildingType"`
}

func getLeases(w http.ResponseWriter, r *http.Request) {
	var leases = []lease{
		{ID: "1", Name: "Condo", BuildingType: "Condominium"},
		{ID: "2", Name: "Apartment", BuildingType: "Apartment"},
		{ID: "3", Name: "House", BuildingType: "House"},
	}

	jsonResponse, jsonError := json.Marshal(leases)
	if jsonError != nil {
		fmt.Println(jsonError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
