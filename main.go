package wordelement

import (
	"log"
	"github.com/bonggeek/element"
	"net/http"
	"encoding/json"
	"strings"
)

type ElementResponse struct{
	Word string
	Elements []element.Element
}

func handler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
    words := q.Get("words")
    if len(words) == 0 {
    	w.WriteHeader(http.StatusBadRequest)
    	w.Write([]byte("Expected query parameters ?words=<words>"))
    	log.Println("No words specified")
    	return
    }

    i := 0
	for _, word := range strings.Split(words, " "){
		i++
		log.Println("Word: ", i, word)
		res := element.GetElements().GetElementsForWord(word)
		elementRes := ElementResponse {Word: word, Elements: res}

		j, _ := json.MarshalIndent(elementRes, "", "    ")
		log.Println("Found elements", len(res))

		w.Header()
		w.Write(j)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
