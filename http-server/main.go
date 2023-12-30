package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Person struct {
	Data interface{} `json:"Data,omitempty"`
	Name string      `json:"Name"`
}

func sendJSONResponse(w http.ResponseWriter, response Person) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error Marshalling json response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	person := Person{Name: "EnTRoPY"}

	sendJSONResponse(w, person)
}

func main() {
	http.Handle("/foo", http.HandlerFunc(fooHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
