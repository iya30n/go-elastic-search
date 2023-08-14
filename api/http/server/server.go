package http

import (
	"log"
	"net/http"
)

func Start() {
    http.HandleFunc("/news", func(w http.ResponseWriter, r *http.Request) {})

    log.Fatalln(http.ListenAndServe(":2000", nil))
}
