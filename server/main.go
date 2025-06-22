// the main server
package main

import (
	"embed"
	"example.com/eaveschallenge/handlers"
	"fmt"
	"log"
	"net/http"
)


const server = "127.0.0.1:5050"

// embed static files into the binary

//go:embed templates/* static/* files/*
var content embed.FS

func main() {
	// serve css file
	http.Handle("/static/", http.FileServer(http.FS(content)))

	// pass embedded FS to handlers package
	handlers.SetFS(content)

	// handlers
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/bot-login", handlers.BotLoginHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/welcome", handlers.WelcomeHandler)
	http.HandleFunc("/download", handlers.DownloadHandler)
	http.HandleFunc("/secret", handlers.FTPHandler)
	http.HandleFunc("/script-output", handlers.ScriptHandler)


	// start server on localhost
	fmt.Println("Server started at ", server)

	log.Fatal(http.ListenAndServe(server, nil))

}
