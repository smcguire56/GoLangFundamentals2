package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// https://gobyexample.com/structs
type msg struct {
	Header  string
	Guess   int
	Message string
}

// https://golang.org/pkg/net/http/ 
func printString(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}


func guessGame(w http.ResponseWriter, r *http.Request) {
	newRandomNum := getRandomNumber(1, 20)

	// https://golang.org/src/net/http/cookie.go
	cookie, err := r.Cookie("target")

	if err != nil {
		cookie = &http.Cookie{Name: "target",
			Value:   strconv.Itoa(newRandomNum), // https://golang.org/pkg/strconv/
			Expires: time.Now().Add(72 * time.Hour),
		}

		http.SetCookie(w, cookie)
	}

	// get the input from guess
	guess, _ := strconv.Atoi(r.FormValue("guess"))
	messageStruct := &msg{Header: "Guess a number between 1 and 20", Guess: guess}
	target, _ := strconv.Atoi(cookie.Value)

	// compare the values of the guess and the random number
	// if they are equal print you got it  
	if target == guess {
		cookie = &http.Cookie{
			Name:    "target",
			Value:   strconv.Itoa(newRandomNum),
			Expires: time.Now().Add(365 * 24 * time.Hour),
		}
		// generate a new number and save as the cookie
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
// generate random numbers
// source http://golangcookbook.blogspot.ie/2012/11/generate-random-number-in-given-range.html
func getRandomNumber(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func main() {
	http.HandleFunc("/", printString)
	http.HandleFunc("/guess", guessGame)
	http.ListenAndServe(":8080", nil)
}
