package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		w.Write([]byte("Home!"))
	default:
		w.Write([]byte("method not recognized"))
	}
}

func main() {
	http.HandleFunc("/", Home)
	fmt.Println("Starting server at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
