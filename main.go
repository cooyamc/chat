package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	//r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/hello", HelloHandler)
	http.Handle("/", r)
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
