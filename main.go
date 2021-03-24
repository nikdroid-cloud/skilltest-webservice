package main

import (
	"net/http"
	"nikdroid-cloud/skillTest/handlers"
)

func main() {
	serverMux := http.NewServeMux()
	helloHandler := handlers.NewHello()
	jsonHandler := handlers.NewIndentJSON()

	serverMux.Handle("/", helloHandler)
	serverMux.Handle("/jsonify", jsonHandler)

	server := http.Server{
		Addr:    ":5000",
		Handler: serverMux,
	}
	server.ListenAndServe()
}
