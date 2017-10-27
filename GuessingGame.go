package main

import (
	"html/template"
	"net/http"
	"strconv"
	"time"
	"math/rand"
)

// https://gobyexample.com/structs
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
	myrand := getRandomNumber(1, 20)

	// Try to read the cookie.
	cookie, err := r.Cookie("target")

	if err != nil {	
		cookie = &http.Cookie{ Name: "target", Value: strconv.Itoa(myrand), Expires: time.Now().Add(72 * time.Hour),
		}

		http.SetCookie(w,cookie)
	}

	guess, _ := strconv.Atoi(r.FormValue("guess"))
	messageStruct := &msg{Title: "Guess a number between 1 and 20", Guess: guess}

	target, _ := strconv.Atoi(cookie.Value)

	if target == guess {
		// correct
		cookie = &http.Cookie{
			Name: "target",
			Value: strconv.Itoa(myrand),
			Expires: time.Now().Add(72 * time.Hour),
		}

		http.SetCookie(w, cookie)
	}

	t, _ := template.ParseFiles("guess.tmpl")
	t.Execute(w, messageStruct)
}

func getRandomNumber(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}


func main() {

	http.HandleFunc("/", printString)
	http.HandleFunc("/guess", guessGame)
	http.ListenAndServe(":8080", nil)
}
