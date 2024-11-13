package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)

var users = map[string]string{"admin": "admin", "user": "password"}
var sessions = map[string]string{"token": "use", "password": "correct"}

func enableCors(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        // Handle preflight requests
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}

// func GetUser(r *http.Request) (*User, error) {
func GetUser(r *http.Request) (string, error) {

	conn, err := pgx.Connect(context.Background(), "postgres://ahmad@127.0.0.1:5432/gda")
	if err != nil {
		fmt.Println(err)
		return "", errors.New("Connection failed")
	}

	defer conn.Close(context.Background())

	var greeting string
	err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)

	fmt.Println(greeting)

	// Extract authorization header
	token := r.Header.Get("Authorization")

	// Change to get from database later
	session := sessions[token]

	// check if session is invalid
	if session == "" {
		return "", errors.New("Invalid session")
	}

	return "", nil
}

func chat(w http.ResponseWriter, req *http.Request) {

	// For now this is just an authenticated endpoint
	_, err := GetUser(req)

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

	GetUser(req)

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
	// Later make this create a user object in database
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
		fmt.Println(username, password)
		// Generate session key
		// send it

		key := "sessionkey"
		sessions[key] = "session" // make this a session object. Find out what it needs to track
		fmt.Fprintf(w, key)

		fmt.Println(key)

		w.WriteHeader(http.StatusOK)
		return
	}

}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /chat", chat)
	mux.HandleFunc("POST /register", register)
	mux.HandleFunc("POST /login", login)

	log.Println("Starting server on port 80")

	err := http.ListenAndServe(":80", enableCors(mux))

	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}

}
