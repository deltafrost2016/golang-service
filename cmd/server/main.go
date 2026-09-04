package main

import (
	"log"
	"net/http"

	"crud-app/internal/router"
)

func main() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router.New(),
	}

	log.Println("listening on", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
