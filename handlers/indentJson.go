package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type IndentJSON struct {
}

func NewIndentJSON() *IndentJSON {
	return &IndentJSON{}
}

func (i *IndentJSON) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m := map[string]string{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	err = json.Unmarshal([]byte(body), &m)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	//var data []byte
	_, err = json.MarshalIndent(m, "", "\t")
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintln(w, m)
}
