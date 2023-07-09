package main

import (
	"log"
	handlers "mdl/lib/handlers"
	"net/http"
)

func main() {

	routes := map[string]func(http.ResponseWriter, *http.Request){
		"/":           handlers.IndexHandler,
		"/about":      handlers.AboutHandler,
		"/authors":    handlers.AuthorHandler,
		"/contacts":   handlers.ContactHandler,
		"/books":      handlers.BooksHandler,
		"/book":       handlers.BookHandler,
		"/citations":  handlers.CitationHandler,
		"/conditions": handlers.ConditionHandler,
	}

	for path, handler := range routes {
		http.HandleFunc(path, handler)
	}

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	log.Println("Server Started at Port:8080")
	log.Println("http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
