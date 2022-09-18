package main

import (
	"encoding/json"
	//"fmt"
	"log"
	"net/http"

	"github.com/onceuponworld/ouw-sdk"
)


func kingdomHandler(w http.ResponseWriter, r *http.Request) {

	switch(r.Method) {
	case http.MethodPost:

		profile := r.FormValue(API_PARAM_PROFILE)

		k := ouwsdk.Kingdom{}

		err := json.Unmarshal([]byte(profile), &k)

		if err != nil {
			log.Println(err)
		} else {

			ouwsdk.SetAdd(ouwsdk.KEY_KINGDOMS, k.Name)

 			ouwsdk.KingdomAdd(k)

		}

	case http.MethodGet:


	default:
		log.Println(ERR_UNSUPPORTED_METHOD)
	}
} // kingdomHandler
