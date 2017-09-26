package main

import (
	"io/ioutil"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/WedgeNix/CubbyChaser/session"
	"github.com/mrmiguu/Loading"
	"github.com/mrmiguu/rest"
)

var (
	sessQ = session.NewQueue()
)

func init() {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "5000"
	}

	http.Handle("/", http.FileServer(http.Dir("client")))
	http.HandleFunc("/createSession", createSession)

	rest.Connect(":" + port)
}

func createSession(w http.ResponseWriter, r *http.Request) {
	done := load.New("creating session")

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	done <- false

	id, ordNums, err := session.ParseIDAndOrderNumbers(string(b))
	if err != nil {
		http.Error(w, "unable to parse html", http.StatusUnprocessableEntity)
		return
	}
	done <- false

	http.Error(w, http.StatusText(http.StatusOK), http.StatusOK)
	done <- false

	go func() {
		sess, err := session.New(id, ordNums)
		if err != nil {
			println(err.Error())
			return
		}
		done <- false

		sessQ.Add <- sess
		done <- true
	}()
}

func main() {
	select {}
}
