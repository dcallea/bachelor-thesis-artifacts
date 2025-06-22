/*
this handler sends a HTTP POST request with form data to a server when the
button is clicked
*/
package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

const address = "http://127.0.0.1:5050/login"

var values = url.Values{"username": {"admin"}, "password": {"supersecretpassword832"}}

func BotLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		resp, err := postRequest(address, values)
		if err != nil {
			log.Fatalf("Failed to send POST request: %v", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Failed to read response body: %v", err)
		}

		fmt.Println("Response Status:", resp.Status)
		fmt.Println("Response Body:", string(body))

		w.Write([]byte("The user just logged in!"))

	} else {
		// don't allow requests other than POST
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func postRequest(
	address string,
	values url.Values,
) (*http.Response, error) {
	resp, err := http.PostForm(address, values)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

