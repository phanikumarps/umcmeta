package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/phanikumarps/umc/umc"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/{entity}/{id}", resolver).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))

}
func resolver(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	entity := vars["entity"]
	id := vars["id"]
	switch entity {
	case "account":
		resp := umc.GetAccount(&id)
		j, err := json.Marshal(resp)
		if err != nil {
			log.Print(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
	default:
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Hello, World!"))
	}
}
