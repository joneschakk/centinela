package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"./ui"
	// _ "./statik"
	// "github.com/rakyll/statik/fs"
)

var c Configuration

func ValidateLoginCredentials(username string, password string, target string) bool {
	if t, ok := c.Targets[target]; ok {
		for i := range c.Users {
			if c.Users[i].Name == username &&
				c.Users[i].Password == password &&
				c.Users[i].Roles.HasAnyRole(t.Roles) {
				return true
			}
		}
	}
	return false
}

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

	username := r.FormValue("username")
	password := r.FormValue("password")
	target := r.FormValue("target")
	domain := r.FormValue("domain")
	fmt.Println("authenticating '" + target + "' with " + "(" + username + ":" + password + ")")
	if target == "" {
		w.WriteHeader(406) // TODO change
		fmt.Fprintf(w, "Authentication target is empty")
		fmt.Fprintf(os.Stderr, "Authentication target is empty")
		return
	}

	if ValidateLoginCredentials(username, password, target) {
		// w.Header().Set("X-Centinela-Redirect-To", target)
		http.SetCookie(w,
			&http.Cookie{
				Name:   "centinela_auth_token",
				Domain: domain,
				Value:  GenerateToken(username, target),
			})

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Authenticated: %s", target)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintf(w, "Authentication failed")
	fmt.Fprintf(os.Stderr, "Authentication failed")
}

func loginPageHandler(w http.ResponseWriter, r *http.Request) {
	// statikFS, err := fs.New()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// indexFile, err := statikFS.Open("index.html")
	target := r.URL.Query().Get("target")
	targetUrl := r.URL.Query().Get("url")
	if target == "" {
		w.WriteHeader(400)
		w.Write([]byte("no target specified"))
		return
	}
	w.WriteHeader(200)
	w.Write([]byte(
		strings.Replace(
			strings.Replace(ui.IndexHTML, "{{AuthTarget}}", target, 1),
			"{{TargetUrl}}", targetUrl, 1)))
}

func authenticatedCheckHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Println("k " + k + " v: " + v[0])
	}
	authFor := r.Header.Get("X-Auth-For")
	authToken := r.Header.Get("X-Auth-Token")
	// a, _ := r.Cookie("")
	// a.Value()
	if authFor == "" {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Authentication target is empty")
		fmt.Fprintf(os.Stderr, "Authentication target is empty")
		return
	}
	if isValidToken(authToken) {
		w.WriteHeader(200)
		fmt.Fprintf(w, "Authorized")
		fmt.Fprintf(os.Stderr, "Authorized")
		return
	}
	w.WriteHeader(401)
	fmt.Fprintf(w, "Not authorized")
	fmt.Fprintf(os.Stderr, "Not authorized")
}

func main() {
	c.Load("config.yml")
	// c.PrintConf()
	// usermap = *c.GetUserMap()
	http.HandleFunc("/authenticate", authenticateHandler)

	http.HandleFunc("/is-authenticated", authenticatedCheckHandler)

	http.HandleFunc("/login", loginPageHandler)
	// http.Handle("/login/", http.StripPrefix("/login/", http.FileServer(statikFS)))

	fmt.Println("starting centinela server @ ", c.GetServerAddress())
	log.Fatal(http.ListenAndServe(c.GetServerAddress(), nil))
}
