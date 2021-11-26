package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Response struct {
	Text           string `json:"Text`
	InstancesOfAnd string `json:"InstancesOfAnd`
}

func GetResponse(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	text := r.URL.Query().Get("x")
	var text_response string
	var instances_response string
	if text == "" {
		text_response = "Invalid String"
		instances_response = "N/A"
	} else {
		text_response = text
		instances_response = fmt.Sprintf("%d", InstancesOfAnd(text))
	}

	response := Response{
		Text:           text_response,
		InstancesOfAnd: instances_response,
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
