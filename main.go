package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Password - needed password
var Password []byte

// LenOfPassword - len of Password
var LenOfPassword int = 10

// UpperCase - password characters
var UpperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// LowCase - password characters
var LowCase = "abcdefghijklmnopqrstuvwxyz"

// DigitalChars - digital characters
var DigitalChars = "1234567890"

// PuncChars - punctuation characters
var PuncChars = "@#^&,."

// RandomPasswordGenerator - return random password
func RandomPasswordGenerator(w http.ResponseWriter, r *http.Request) {
	//get nano second number
	var nanoSecond int64 = time.Now().UnixNano() / 1e6

	// clear history string
	Password = []byte{}

	// refresh new seed
	r1 := rand.New(rand.NewSource(nanoSecond))

	// a counter
	var i int = 0

	// rule 1: first character must be upper case letter
	// rule 2: last character must be low case letter
	for i = 0; i < LenOfPassword; i++ {
		remainder1 := r1.Intn(1000) % 4
		if i == 0 {
			remainder1 = 0
		}
		if i == LenOfPassword-1 {
			remainder1 = 1
		}
		switch remainder1 {
		case 0:
			remainder2 := r1.Intn(1000) % len(UpperCase)
			Password = append(Password, []byte(UpperCase)[remainder2])
		case 1:
			remainder2 := r1.Intn(1000) % len(LowCase)
			Password = append(Password, []byte(LowCase)[remainder2])
		case 2:
			remainder2 := r1.Intn(1000) % len(DigitalChars)
			Password = append(Password, []byte(DigitalChars)[remainder2])
		case 3:
			remainder2 := r1.Intn(1000) % len(PuncChars)
			Password = append(Password, []byte(PuncChars)[remainder2])
		}
	}
	fmt.Fprintf(w, string(Password))
}

func main() {

	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/", RandomPasswordGenerator)
	http.Handle("/", gorillaRoute)
	http.ListenAndServe(":8080", nil)
}
