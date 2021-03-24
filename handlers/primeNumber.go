package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type PrimeNumber struct {
}

func NewPrimeNumber() *PrimeNumber {
	return &PrimeNumber{}
}

func (pm *PrimeNumber) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	num := r.URL.String()[1:]

	numSpecified, err := strconv.Atoi(num)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "couldn't be able to get the number")
		return
	}

	message := map[string]string{
		"message": isPrimeNumber(numSpecified),
	}

	toReturned, err := json.MarshalIndent(message, "", "")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	fmt.Fprint(w, string(toReturned))
}

func isPrimeNumber(n int) string {
	if n == 1 {
		return "False"
	}
	for i := 2; i <= n/2; i++ {
		// condition for non-prime
		if n%i == 0 {
			return "False"
		}
	}
	return "True"
}
