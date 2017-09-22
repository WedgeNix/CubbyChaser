package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/WedgeNix/CubbyChaser-shared"
	"github.com/WedgeNix/CubbyChaser/session"
	"github.com/mrmiguu/Loading"
	"github.com/mrmiguu/rest"
)

var (
	newSession = make(chan shared.Session, 10)
	sessions   = map[int]shared.Session{}
)

func main() {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "5000"
	}
	http.Handle("/", http.FileServer(http.Dir("client")))
	http.HandleFunc("/createSession", createSession)
	rest.Connect(":" + port)

	for sess := range newSession {
		go addSession(sess)
	}
}

func createSession(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	id, ordNums, err := session.ParseIDAndOrderNumbers(string(b))
	if err != nil {
		http.Error(w, "unable to parse html", http.StatusUnprocessableEntity)
		return
	}

	http.Error(w, http.StatusText(http.StatusOK), http.StatusOK)

	go func() {
		sess, err := session.New(id, ordNums)
		if err != nil {
			println(err.Error())
			return
		}
		newSession <- sess
	}()
}

func addSession(sess shared.Session) {
	sessions[sess.ID] = sess
	fmt.Println(sessions[sess.ID])

	// BEWARE: multiple connection REST channels made on global line
	_, connection := rest.Bool()
	session, _ := rest.Bytes()
	for {
		done := load.Ing(`connection()`)
		connection()
		done <- true

		session(shared.Stob(sessions[sess.ID]))
	}
}
