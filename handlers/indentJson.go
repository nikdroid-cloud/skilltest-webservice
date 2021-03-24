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
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "couldn't be able to read the body of the request.")
		return
	}

	var dynamicJSON map[string]string
	err = json.Unmarshal(body, &dynamicJSON)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	indentedJSON, err := json.MarshalIndent(dynamicJSON, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	fmt.Fprint(w, addTabEscaped(indentedJSON))
}

func addTabEscaped(data []byte) string {
	str := string(data)

	//remove last curly bracket
	strBeforeEnd := str[:len(str)-1]

	//add tab escaped and curly bracket in the end of the json
	return strBeforeEnd + "\t}"
}
