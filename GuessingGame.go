package main

import (
	"fmt"
  	"net/http"
)

func printString(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "<h1>Guessing Game</h1>")

}

func main() {
	http.HandleFunc("/", printString)
	http.ListenAndServe(":8080", nil)
}
