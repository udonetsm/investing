package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/udonetsm/investing/server/controllers"
)

func StartServer() {
	server := buildServer(":8080")
	fmt.Print("OK\n")
	server.ListenAndServe()

}

func buildServer(addr string) *http.Server {
	fmt.Print("Starting server...")
	return &http.Server{
		Addr:              addr,
		Handler:           router(),
		ReadTimeout:       3 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      3 * time.Second,
		IdleTimeout:       3 * time.Second,
	}
}

func router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/{count-show-read}/{all-new}/{requests-replies}/of/{uid}", controllers.SubOnSomething).Methods(http.MethodPost)
	// {tid} should be current transaction id
	// {requests-replies} should be "requests", or "replies"
	// {uid} should be an exist user id
	router.HandleFunc("/pub/{tid}/to/{requests-replies}/of/{uid}", controllers.PubMessage)
	// Saves transaction with {id} to the database.
	router.HandleFunc("/save/{something}/{id}", controllers.CaptureSomething).Methods(http.MethodPost)
	router.HandleFunc("/info/{about}/{id}", controllers.Info)

	return router
}
