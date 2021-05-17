package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"scolix/cmd/api0/httpHandlers"
	"scolix/pkg/httpmiddleware"
)

func main() {
	router := mux.NewRouter().StrictSlash(false)

	// Use the middlewares
	router.Use(httpmiddleware.EnableCors)
	router.Use(httpmiddleware.SetResponseTypeJSON)
	router.Use(httpmiddleware.RequestLogger)


	router.HandleFunc("/", httpHandlers.Index)

	if len(os.Args) > 1 {
		if os.Args[1] == "-dev" {
			log.Println("The server is running...")
		}
	}

	log.Fatal(http.ListenAndServe(":8080", router))
}