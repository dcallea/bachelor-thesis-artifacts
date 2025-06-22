package handlers

import (
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/html")
		renderTemplate(w, "index.html", nil)
	} else {
		// don't allow requests other than GET
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

