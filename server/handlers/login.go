package handlers

import (
	//"encoding/json"
	"fmt"
	"net/http"
	"github.com/google/uuid"
)


// represent all valid users in a map
var users = map[string]string{
	"admin": "supersecretpassword832",
}

var sessions = map[string]string{}

func valueExists(m map[string]string, value string) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}
	return false
}

// checks if username and password exists and is valid (in users map)
func login(username string, password string) bool {
	storedPassword, ok := users[username]
	// if username does not exist in users
	if !ok {
		return false
	}
	// else return whether password matches with saved password
	return password == storedPassword
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		renderTemplate(w, "login.html", nil)
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if login(username, password) {
			// if user already has session active
			if valueExists(sessions, username) {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, "You are already logged in.")
				return
			}

			// create new session token and store session in session map
			sessionToken := uuid.NewString()
			sessions[sessionToken] = username

			// set the cookie to store session token clientside
			http.SetCookie(w, &http.Cookie{
				Name:	"session_token",
				Value:	sessionToken,
				Secure: false,
				Path: 	"/",
			})

			fmt.Fprintf(w, "Login successful! Welcome, %s\n", username)


		} else {
			fmt.Fprintf(w, "Invalid login credentials. Please try again")
		}
	}
}
