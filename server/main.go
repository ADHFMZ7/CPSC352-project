package main

import (
	"fmt"
	"net/http"
	"log"
)

var users = map[string]string{"admin": "admin", "user": "password"}
var sessions = map[string]string{"token": "session"}

func chat(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func register(w http.ResponseWriter, req *http.Request) {
	// Register a new user

	// Get the username and password from the request
	username := req.FormValue("username")
	password := req.FormValue("password")

	// Check if the username is already taken
	if _, ok := users[username]; ok {
		fmt.Fprintf(w, "Username already taken")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Add the new user to the users map
	users[username] = password
	fmt.Fprintf(w, "User registered successfully")

	log.Printf("Registered user %s with password %s\n", username, password)

	w.WriteHeader(http.StatusOK)
	return
}

func login(w http.ResponseWriter, req *http.Request) {

	username := req.FormValue("username")
	password := req.FormValue("password")

	if users[username] == password {
		// Generate session key
		// send it
		key := "SESSION KEY"
		fmt.Fprintf(w, key)
		w.WriteHeader(http.StatusOK)
		return
	}

	fmt.Println(req)

}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /chat", chat)
	mux.HandleFunc("POST /register", register)
	mux.HandleFunc("POST /login", login)

	log.Println("Starting server on port 80")

	err := http.ListenAndServe(":80", mux)

	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}

}
