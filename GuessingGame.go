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
	Header string
	Guess int
	Message string
}

func printString(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func guessGame(w http.ResponseWriter, r *http.Request) {
	newRandomNum := getRandomNumber(1, 20)

	// https://golang.org/src/net/http/cookie.go
	cookie, err := r.Cookie("target")

	if err != nil {	
		cookie = &http.Cookie{ Name: "target",
							   Value: strconv.Itoa(newRandomNum),				// https://golang.org/pkg/strconv/
							   Expires: time.Now().Add(72 * time.Hour),
		}

		http.SetCookie(w,cookie)
	}

	guess, _ := strconv.Atoi(r.FormValue("guess"))
	messageStruct := &msg{Header: "Guess a number between 1 and 20", Guess: guess}
	target, _ := strconv.Atoi(cookie.Value)

	if target == guess {
		cookie = &http.Cookie{
			Name: "target",
			Value: strconv.Itoa(newRandomNum),
			Expires: time.Now().Add(365 * 24 * time.Hour),
		}

		http.SetCookie(w, cookie)

		messageStruct.Message = "You Guessed Correctly!"
	} else if guess < target {
		messageStruct.Message = "Too low, try again."
	} else if guess > target {
		messageStruct.Message = "Too high, try again"
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
