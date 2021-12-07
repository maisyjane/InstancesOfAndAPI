package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Response struct {
	Text       string `json:"Text`
	Answer     string `json:"answer`
	HttpStatus string `json:"HttpStatus`
}

func OutputResponse(text string) Response {
	var text_response string
	var instances_response string
	var code string
	if text == "" {
		text_response = "Invalid String"
		instances_response = "0"
		code = "404"

	} else {
		text_response = text
		instances_response = fmt.Sprintf("%d", InstancesOfAnd(text))
		code = "200"
	}

	response := Response{
		Text:       text_response,
		Answer:     instances_response,
		HttpStatus: code,
	}

	return response

}
func GetResponse(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	text := r.URL.Query().Get("x")
	response := OutputResponse(text)
	if response.HttpStatus == "200" {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(404)
	}
	json.NewEncoder(w).Encode(response)

}
func main() {
	http.HandleFunc("/", GetResponse)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func InstancesOfAnd(text string) int {
	var counter int = 0
	replacer := strings.NewReplacer(",", "", ".", "", ";", "")
	text = replacer.Replace(text)
	words := strings.Fields(text)
	for i := range words {
		if words[i] == "and" {
			counter++
		}

	}
	return counter
}

//see bookmarks for json
