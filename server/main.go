package main

import (
	"fmt"
	"net/http"
	"log"
	"errors"
)

var users = map[string]string{"admin": "admin", "user": "password"}
var sessions = map[string]string{"token": "use", "password":"correct"}

// func GetUser(r *http.Request) (*User, error) {
func GetUser(r *http.Request) (string, error) {

	token := r.Header.Get("Authorization")

	session := sessions[token]


	if session == "" {
		return "", errors.New("Invalid session")
	}
	
	return "", nil
}


func chat(w http.ResponseWriter, req *http.Request) {
	
	// For now this is just an authenticated endpoint
	_ , err := GetUser(req)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Successfully authenticated user")

}


func create_chat(w http.ResponseWriter, req *http.Request) {

	// Should take in the follwing parameters
	//
	// users: list of user ids for the new chat

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

		key := "sessionkey"
		sessions[key] = "session" // make this a session object. Find out what it needs to track
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
