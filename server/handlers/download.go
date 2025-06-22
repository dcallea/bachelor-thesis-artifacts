package handlers

import (
	"io"
	"net/http"
	"fmt"
)


func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	sessionCookie, err := r.Cookie("session_token")
	if err == http.ErrNoCookie {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "You are not authorized to view this page. Please log in first.")
		return
	} else if err != nil {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal Server Error")
		return
	}

	sessionToken := sessionCookie.Value
	_, exists := sessions[sessionToken]

	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	file, err := fs.Open("files/access_log_anonymized.txt")
	if err != nil {
		http.NotFound(w, r)
		return
	}
	defer file.Close()
		
	w.Header().Set("Content-Disposition", "attachment; filename=access_log_anonymized.txt")
	w.Header().Set("Content-Type", "text/plain")
	io.Copy(w, file)
}
