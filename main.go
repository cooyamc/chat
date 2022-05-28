package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/hello", HelloHandler)
	r.HandleFunc("/messages", MessagesGetHandler).Methods("GET")
	r.HandleFunc("/messages", MessagesPostHandler).Methods("POST")
	http.Handle("/", r)
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

type Message struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

var messages = []Message{}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("frontend.html")
	if err != nil {
		log.Printf("failed to read frontend: %s", err)
	}
	w.Write(data)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World!<h1>"))
}

func MessagesGetHandler(w http.ResponseWriter, r *http.Request) {
	j, err := json.Marshal(messages)
	if err != nil {
		log.Printf("failed to marshal messages: %s", err)
		return
	}

	w.Write(j)
}

func MessagesPostHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("failed to read the request body: %s", err)
		return
	}

	var m Message
	err = json.Unmarshal(body, &m)
	if err != nil {
		log.Printf("failed to unmarshal the body: %s", err)
		return
	}

	messages = append(messages, m)
	fmt.Fprintln(w, "{\"status\": \"ok\"}")
}
