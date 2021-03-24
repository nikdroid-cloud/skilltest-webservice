package handlers

import (
	"fmt"
	"net/http"
	"time"
)

type Hello struct {
}

func NewHello() *Hello {
	return &Hello{}
}

func (i *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s\n", time.Now().Format(time.ANSIC))
}
