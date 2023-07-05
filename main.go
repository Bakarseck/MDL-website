package main

import (
	"log"
	handlers "mdl/lib/handlers"
	"net/http"
)


func main() {

	routes := map[string]func(http.ResponseWriter, *http.Request){
		"/": handlers.IndexHandler,
	}

	for path, handler := range routes {
		http.HandleFunc(path, handler)
	}

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	log.Println("Server Started at Port:8080")
	log.Println("http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
