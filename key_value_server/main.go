package main

import (
	"sync"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Value struct {
	ID        string   `json:"id,omitempty`
	Value     string   `json:"value,omitempty`
	Type      string   `json:"type,omitempty"`
}

var m map[string]Value
var mutex = &sync.Mutex{}

func GetValueEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	mutex.Lock()
	defer mutex.Unlock()
	var item = m[params["id"]]
	json.NewEncoder(w).Encode(item)
}
func GetValuesEndpoint(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	keys := [] Value {}
	for _, item := range m {
       keys = append(keys, item)
	}
	json.NewEncoder(w).Encode(keys)
}
func CreateValueEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var value Value
	_ = json.NewDecoder(r.Body).Decode(&value)
	value.ID = params["id"]

	mutex.Lock()
	defer mutex.Unlock()
	m[value.ID] = value
	json.NewEncoder(w).Encode(value)
}
func DeleteValueEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
    mutex.Lock()
	defer mutex.Unlock()
	value := m[params["id"]]
    delete(m, params["id"])
    json.NewEncoder(w).Encode(value)
}

func main() {
	m = make(map[string]Value)
	router := mux.NewRouter()

	m["1"] = Value{ID: "1", Value: "Carlos", Type: "string"}
	m["2"] = Value{ID: "2", Value: "1.2000", Type: "float"}

	router.HandleFunc("/value", GetValuesEndpoint).Methods("GET")
	router.HandleFunc("/value/{id}", GetValueEndpoint).Methods("GET")
	router.HandleFunc("/value/{id}", CreateValueEndpoint).Methods("POST")
	router.HandleFunc("/value/{id}", DeleteValueEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
