package main

import (
	"net/http"

	"github.com/nikdroid-cloud/skilltest-webservice/handlers"
)

func main() {

	serveMux := http.NewServeMux()
	helloHandler := handlers.NewHello()
	indentJSONHandler := handlers.NewIndentJSON()
	primeNumberHandler := handlers.NewPrimeNumber()

	serveMux.Handle("/", helloHandler)
	serveMux.Handle("/stringify", indentJSONHandler)
	serveMux.Handle("/check_number/", http.StripPrefix("/check_number", primeNumberHandler))

	server := http.Server{
		Addr:    ":5000",
		Handler: serveMux,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
