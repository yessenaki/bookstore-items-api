package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func StartApp() {
	mapUrls()

	server := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
	}

	log.Fatal(server.ListenAndServe())
}
