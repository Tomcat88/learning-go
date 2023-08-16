package main

import (
	// "encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(w io.Writer, name string) {
	/* jsonRespose := make(map[string]string)
	jsonRespose["respose"] = "Hello, " + name
	j, err := json.Marshal(jsonRespose)
	if err != nil {
		fmt.Printf("Error %s", err)
	}
	fmt.Fprintf(w, string(j)) */
	fmt.Fprintf(w, "Hello, " + name)
}

func GreeterHandler(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Query()
	who := val.Get("who")
	if who == "" {
		who = "no one"
	}
	Greet(w, who)
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(GreeterHandler)))
}
