package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// after running this script, use following command:
// curl -XGET -H "Content-type: application/json" -d '{"Firstname":"Professor", "Lastname": "Sensei", "Age":"Eternal"}' 'localhost:8080/decode'

type Sensei struct {
	Firstname string
	Lastname  string
	Age       string // in real life this would be an int lol
	Proverbs  []string
}

func main() {
	// encode
	http.HandleFunc("/encode", enco)
	// decode
	http.HandleFunc("/decode", deeco)
	// Listen and serve
	http.ListenAndServe(":8080", nil)
}

// encode
func enco(w http.ResponseWriter, r *http.Request) {
	s1 := Sensei{
		Firstname: "Professor",
		Lastname:  "Sensei",
		Age:       "forever",
		Proverbs:  []string{"Professor Sensei, bearer of light", "Professor Sensei is", "What can you do for your Sensei?"},
	}

	err := json.NewEncoder(w).Encode(s1)
	if err != nil {
		log.Println("Encoded bad dad for Sensei!", err)
	}
}

// decode
func deeco(w http.ResponseWriter, r *http.Request) {
	var s1 Sensei
	err := json.NewDecoder(r.Body).Decode(&s1)
	if err != nil {
		log.Println("Decoded bad data for Sensei!", err)
	}

	log.Println(s1.Firstname, " ", s1.Lastname, "is ", s1.Age)
}
