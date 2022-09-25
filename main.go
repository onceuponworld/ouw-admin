package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/onceuponworld/ouw-sdk"
)


func initRoutes() *mux.Router {

	router := mux.NewRouter()
	
	router.HandleFunc("/api/kingdoms", kingdomHandler)
	router.HandleFunc("/api/kingdoms/{kid:[a-z]+}", kingdomHandler)
	router.HandleFunc("/api/kingdoms/{kid:[a-z]+}/municipals", municipalHandler)
	router.HandleFunc("/api/kingdoms/{kid:[a-z]+}/municipals/{mid:[a-z]+}", municipalHandler)
	router.HandleFunc("/api/npcs", npcHandler)

  return router

} // initRoutes


func main() {

	ouwsdk.Init(true)

	address := ouwsdk.GetAddr()

	routes := initRoutes()

	fmt.Printf("Starting %s on %s...\n", APP_NAME, address)

	log.Fatal(http.ListenAndServe(address, routes))

} // main
