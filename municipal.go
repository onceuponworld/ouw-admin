package main

import (
	//"encoding/json"
	//"fmt"
	"log"
  "net/http"

	"github.com/gorilla/mux"
	"github.com/onceuponworld/ouw-sdk"
)


func municipalHandler(w http.ResponseWriter, r *http.Request) {

	m := mux.Vars(r)

	kid := m[QUERY_KID]

	if ouwsdk.SetExist(ouwsdk.KingdomsKey(), kid) {

		switch(r.Method) {
		case http.MethodPost:

			name := r.FormValue(API_PARAM_NAME)

			if ouwsdk.CheckStr(name, DEFAULT_NAME_LENGTH_MIN) {

				key := ouwsdk.MunicipalsKey(kid)

				ouwsdk.SetAdd(key, name)

			} else {
				w.WriteHeader(http.StatusBadRequest)
			}

		case http.MethodPut:
			
			k := ouwsdk.MunicipalsKey(kid)
			
			log.Println(k)
	
			/*
			err := json.Unmarshal([]byte(profile), &k)
	
			if err != nil {
				log.Println(err)
			} else {
	
				//ouwsdk.KingdomAdd(k)
	
			}
			*/
	
		case http.MethodGet:
	
	
		default:
			log.Println(ERR_UNSUPPORTED_METHOD)
		}
	
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

} // municipalHandler
