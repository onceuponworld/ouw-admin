package main

import (
	"encoding/json"
	//"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/onceuponworld/ouw-sdk"
)


func kingdomHandler(w http.ResponseWriter, r *http.Request) {

	switch(r.Method) {
	case http.MethodPost:

		name := r.FormValue(API_PARAM_NAME)

		ouwsdk.SetAdd(ouwsdk.KingdomsKey(), name)

		k := ouwsdk.Kingdom{}

		k.Name = name

		ouwsdk.KingdomAdd(k)

	case http.MethodPut:

		vars := mux.Vars(r)

		kid := vars[QUERY_KID]

		if ouwsdk.SetExist(ouwsdk.KingdomsKey(), kid) {

			k := ouwsdk.KingdomGet(kid)

			log.Println(k)

		} else {
			w.WriteHeader(http.StatusNotFound)
		}

	case http.MethodGet:

		vars := mux.Vars(r)

		kid := vars[QUERY_KID]

		if ouwsdk.SetExist(ouwsdk.KingdomsKey(), kid) {

			k := ouwsdk.KingdomGet(kid)

			log.Println(k)

			j, err := json.Marshal(k)

			if err != nil {
				log.Println(err)
			} else {
				w.Write(j)
			}

		} else {
			w.WriteHeader(http.StatusNotFound)
		}


	default:
		log.Println(ERR_UNSUPPORTED_METHOD)
	}

} // kingdomHandler
