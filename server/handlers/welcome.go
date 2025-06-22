package handlers

import (
	"fmt"
	"net/http"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	sessionCookie, err := r.Cookie("session_token")
	if err != nil {
		// if session_token is not set, return unauthorized status
		if err == http.ErrNoCookie {
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "You are not authorized to view this page. Please log in first.")
			return
		}
		// for any other error, return bad request status
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "Bad Request.")
		return
	}
	sessionToken := sessionCookie.Value
	// check if session exists
	_, exists := sessions[sessionToken]
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "You are not authorized to view this page. Please log in first.")
		return
	}

	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "You are authorized as %s.\n", sessions[sessionToken])
		renderTemplate(w, "welcome.html", nil)

	} else {
		// need to change to download stuff
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

