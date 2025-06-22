package handlers

import (
	"net/http"
)

func FTPHandler(w  http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/html")
		renderTemplate(w, "ftp.html", nil)
	}
}
