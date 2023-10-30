package main

import (
	"log"
	"net/http"

	"github.com/yudiagonal/go-api-simple-checkout-products/database"
	"github.com/yudiagonal/go-api-simple-checkout-products/internal/factory"
)

func main() {

	db, err := database.Connect()

	if err != nil {
		log.Fatalln(err)
	}
	mux := http.NewServeMux()

	factory.RegisterHandlers(mux, db)

	port := ":8080"
	log.Println("Listening Server On Port " + port)

	err = http.ListenAndServe(port, mux)

	if err != nil {
		log.Fatalln("Failed to start Server")
	}

}
