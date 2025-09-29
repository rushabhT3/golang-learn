package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// mux: popular router for golang: uber use it
	router := mux.NewRouter()

	router.HandleFunc("/books/{title}/pages/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]
		fmt.Fprintf(w, "Book: %s\nPage: %s\n", title, page)
	})

	// CreateBook, ReadBook, UpdateBook, and DeleteBook are functions; that must be define in the codebase
	/*
		router.HandleFunc("/books/{title}", CreateBook).Methods("POST")
		router.HandleFunc("/books/{title}", ReadBook).Methods("GET")
		router.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
		router.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")
	*/

	fmt.Println("Server is running on port 80")
	// port 80: don't have to write that in the URL section like localhost:80; writing localhost only also works
	http.ListenAndServe(":80", router)
}
