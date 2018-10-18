package main

import (
	"fmt"
	"log"
	"net/http"

	_ "./statik"
	"github.com/rakyll/statik/fs"
)

func authenticateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method == "OPTIONS" { // CORS preflight
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.WriteHeader(200)
		return
	}

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		return
	}

	authFor := r.Header.Get("X-Auth-For")
	if authFor == "" {
		w.WriteHeader(406) // TODO change
		fmt.Fprintf(w, "Authentication target is empty")
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println("authenticating ..... | " + username + ":" + password)

	if username == "u" && password == "p" {
		http.Redirect(w, r, authFor, http.StatusTemporaryRedirect)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "Authentication failed")
}

func loginPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func authenticatedCheckHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Println("k " + k + " v: " + v[0])
	}
	authFor := r.Header.Get("X-Auth-For")
	authToken := r.Header.Get("X-Auth-Token")
	if authFor == "" {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Authentication target is empty")
		return
	}
	if authToken == "foo" {
		w.WriteHeader(200)
		fmt.Fprintf(w, "Authorized")
		return
	}
	w.WriteHeader(401)
	fmt.Fprintf(w, "Not authorized")
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
	log.Fatal(http.ListenAndServe(":6969", nil))
}
