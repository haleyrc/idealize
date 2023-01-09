package main

import (
	"log"
	"net/http"
)

func main() {
	r := newRouter()

	r.HandleFunc("/GetMessage", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":"Hello!"}`))
	})

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
