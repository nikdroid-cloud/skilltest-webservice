package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/nikdroid-cloud/skillTest/handlers"
)

func TestHelloHandler(t *testing.T) {
	serveMux := http.NewServeMux()
	helloHandler := handlers.NewHello()
	serveMux.Handle("/", helloHandler)
	server := httptest.NewServer(serveMux)

	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
	}
	expected := "Hello World!\n"
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	actual := string(body)
	if expected != actual {
		t.Errorf("Expected %v got %v", expected, actual)
	}
}

func TestIndentJSONHandler(t *testing.T) {
	serveMux := http.NewServeMux()
	iJSONHandler := handlers.NewIndentJSON()
	serveMux.Handle("/stringify", iJSONHandler)
	server := httptest.NewServer(serveMux)

	defer server.Close()

	requestBody, err := json.Marshal(map[string]string{
		"name":    "Mario",
		"surname": "Rossi",
	})
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Post(server.URL+"/stringify", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
	}
	expected := "{\n \"name\": \"Mario\",\n \"surname\": \"Rossi\"\n\t}"
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	actual := string(body)
	if expected != actual {
		t.Errorf("Expected \n%v \ngot \n%v", expected, actual)
	}
}

func TestPrimeNumberHandler(t *testing.T) {
	serveMux := http.NewServeMux()
	primeNumberHandler := handlers.NewPrimeNumber()
	serveMux.Handle("/check_number/", http.StripPrefix("/check_number", primeNumberHandler))
	server := httptest.NewServer(serveMux)

	defer server.Close()

	primeNumbers := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

	for _, v := range primeNumbers {
		resp, err := http.Post(server.URL+"/check_number/"+strconv.Itoa(v), "application/json", nil)
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != 200 {
			t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
		}
		expected := "{\n\"message\": \"True\"\n}"
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		actual := string(body)
		if expected != actual {
			t.Errorf("Expected \n%v \ngot \n%v", expected, actual)
		}
	}
}
