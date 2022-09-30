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

			kingdom := r.FormValue(API_PARAM_KINGDOM)

			log.Println(kingdom)

			if ouwsdk.CheckStr(kingdom, ouwsdk.DEFAULT_NAME_LENGTH_MIN) {

				kingdomMod := ouwsdk.Kingdom{}

				err := json.Unmarshal([]byte(kingdom), &kingdomMod)
	
				if err != nil {
					log.Println(err)
				} else {
					log.Println(kingdomMod)
					ouwsdk.KingdomUpdate(kid, kingdomMod)
				}
	
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}


		} else {
			w.WriteHeader(http.StatusNotFound)
		}

	case http.MethodGet:

		vars := mux.Vars(r)

		kid := vars[QUERY_KID]

		if len(kid) == 0 {

			kingdoms := ouwsdk.KingdomGetAll()

			buf := ouwsdk.ToJson(kingdoms)

			w.Write(buf)
		
		} else {

			if ouwsdk.SetExist(ouwsdk.KingdomsKey(), kid) {

				k := ouwsdk.KingdomGet(kid)
	
				buf := ouwsdk.ToJson(k)

				w.Write(buf)
	
			} else {
				w.WriteHeader(http.StatusNotFound)
			}

		}

	default:
		log.Println(ERR_UNSUPPORTED_METHOD)
	}

} // kingdomHandler
