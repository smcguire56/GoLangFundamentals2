package main

import (
	"html/template"
	"net/http"
	"strconv"
)

type msg struct {
	Title string
	Guess int
}

func printString(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html") // allows browser to render html tags
	//fmt.Fprint(w, "<h1>Guessing Game</h1>")
	http.ServeFile(w, r, "index.html")
}
func guessGame(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set("Content-Type", "text/html") // allows browser to render html tags
	//fmt.Fprint(w, "<h1>Guessing Game</h1>")
	//http.ServeFile(w, r, "guess.html")

	guess, _ := strconv.Atoi(r.FormValue("guess"))
	messageStruct := &msg{Title: "Guess a number between 1 and 20", Guess: guess}

	t, _ := template.ParseFiles("guess.tmpl")
	t.Execute(w, messageStruct)
}

func main() {

	http.HandleFunc("/", printString)
	http.HandleFunc("/guess", guessGame)
	http.ListenAndServe(":8080", nil)
}
