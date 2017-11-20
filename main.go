package main

import (
	"log"
	"github.com/bonggeek/element"
	"net/http"
	"encoding/json"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	words := q.Get("words")
	for _, word := range strings.Split(words, " "){
		log.Println(word)
		res := element.GetElements().GetElementsForWord(word)

		j, _ := json.MarshalIndent(res, "", "    ")

		log.Println(string(j))
		w.Header()
		w.Write(j)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
