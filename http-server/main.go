package main

import (
	"encoding/json"
	"fmt"
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

func createPerson(w http.ResponseWriter, r *http.Request) {
	var person *Person

	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		panic(err)
	}

	sendJSONResponse(w, *person)
}

func getPerson(w http.ResponseWriter) {
	person := Person{Name: "EnTRoPY"}

	sendJSONResponse(w, person)
}

func personHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Println("GET")
		getPerson(w)
	case http.MethodPost:
		fmt.Println("POST")
		createPerson(w, r)
	}
}

func main() {
	http.Handle("/foo", http.HandlerFunc(personHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
