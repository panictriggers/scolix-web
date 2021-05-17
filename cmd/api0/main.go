package main

import (
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"scolix/cmd/api0/httpHandlers"
	"scolix/pkg/httpmiddleware"
)

func main() {
	var dev = flag.Bool("dev", false, "enable dev debug messages")
	flag.Parse()
	router := mux.NewRouter().StrictSlash(false)

	// Use the middlewares
	router.Use(httpmiddleware.EnableCors)
	router.Use(httpmiddleware.SetResponseTypeJSON)

	// Dev middlewares
	if *dev == true {
		router.Use(httpmiddleware.RequestLogger)
	}


	router.HandleFunc("/", httpHandlers.Index)

	if *dev == true {
		log.Println("The server is running")
	}

	log.Fatal(http.ListenAndServe(":8080", router))
}