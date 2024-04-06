package main

import (
	"log"
	"net/http"
	"shoeShop/db"
	"time"
)

func main() {
	m := http.NewServeMux()

	m.Handle("/", http.HandlerFunc(mainHandle))

	server := http.Server{
		Addr:         ":8008",
		Handler:      m,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	db.Connect()
	db.Migrate()

	log.Println("Listening localhost:8008")

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
