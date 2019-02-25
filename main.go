package main

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", Ping)
	log.Info("starting golang server on port 8080")
	http.ListenAndServe(":8080", mux)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ping"))
}
