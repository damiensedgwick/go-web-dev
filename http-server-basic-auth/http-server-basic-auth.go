package main

import (
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
)

const (
	HOST     = "localhost"
	PORT     = "8080"
	USER     = "admin"
	PASSWORD = "password"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func BasicAuth(h http.HandlerFunc, realm string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || subtle.ConstantTimeCompare([]byte(user),
			[]byte(USER)) != 1 || subtle.ConstantTimeCompare([]byte(pass),
			[]byte(PASSWORD)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("You are not authorised to access this application.\n"))
			return
		}
		h(w, r)
	}
}

func main() {
	http.HandleFunc("/", BasicAuth(helloWorld, "Please enter your username and password"))
	err := http.ListenAndServe(HOST+":"+PORT, nil)
	if err != nil {
		log.Fatalln("Error starting http server : ", err)
		return
	}
}
