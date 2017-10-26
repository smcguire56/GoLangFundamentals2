package main

import (
  	"net/http"
)

func printString(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type","text/html");// allows browser to render html tags
	//fmt.Fprint(w, "<h1>Guessing Game</h1>")
	http.ServeFile(w, r,"index.html")
}
func guessGame(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type","text/html");// allows browser to render html tags
	//fmt.Fprint(w, "<h1>Guessing Game</h1>")
	http.ServeFile(w, r,"guess.html")
}

func main() {
	http.HandleFunc("/", printString)
	http.HandleFunc("/guess", guessGame)
	
	
	
	http.ListenAndServe(":8080", nil)
}
