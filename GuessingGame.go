package main

import (
  	"net/http"
)

func printString(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type","text/html");// allows browser to render html tags
	//fmt.Fprint(w, "<h1>Guessing Game</h1>")
	http.ServeFile(w, r,"index.html")
}

func main() {
	http.HandleFunc("/", printString)
	http.ListenAndServe(":8080", nil)
}
