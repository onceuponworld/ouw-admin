package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/onceuponworld/ouw-sdk"
)


func main() {

	ouwsdk.Init(true)

	address := ouwsdk.GetAddr()

	fmt.Printf("Starting %s on %s...\n", APP_NAME, address)

	http.HandleFunc("/api/kingdoms", kingdomHandler)
	http.HandleFunc("/api/npcs", npcHandler)

	log.Fatal(http.ListenAndServe(address, nil))

} // main
