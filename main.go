package main

import (
	"fmt"
	"log"
	"net/http"

	_ "./statik"
	"github.com/rakyll/statik/fs"
)

func authenticateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func loginPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func authenticatedCheckHandler(w http.ResponseWriter, r *http.Request) {
	// r.Response.StatusCode = 500
	w.WriteHeader(200)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/authenticate", authenticateHandler)

	http.HandleFunc("/is-authenticated", authenticatedCheckHandler)

	statikFS, err := fs.New()

	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/login/", http.StripPrefix("/login/", http.FileServer(statikFS)))

	fmt.Println("starting centinela server...")
	log.Fatal(http.ListenAndServe(":9999", nil))
}
