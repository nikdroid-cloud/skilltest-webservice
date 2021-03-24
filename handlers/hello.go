package handlers

import (
	"fmt"
	"net/http"
)

type Hello struct {
}

func NewHello() *Hello {
	return &Hello{}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
