package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/WedgeNix/CubbyChaser-shared"
	"github.com/mrmiguu/Loading"
	"github.com/mrmiguu/rest"
)

var (
	sessions = map[int]shared.Session{}
)

func main() {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "5000"
	}
	http.Handle("/", http.FileServer(http.Dir("client")))
	rest.Connect(":" + port)

	// TODO: pipe real HTML document here
	b, err := ioutil.ReadFile("test2.html")
	sess, err := sessionStart(string(b))
	if err != nil {
		panic(err)
	}
	sessions[sess.ID] = sess
	fmt.Println(sessions[sess.ID])

	_, connection := rest.Bool()
	session, _ := rest.Bytes()
	for {
		done := load.Ing(`connection()`)
		connection()
		done <- true

		b, _ = shared.Stob(sessions[sess.ID])
		session(b)
	}
}

var (
	sessionLine          = regexp.MustCompile(`(?i)>[^<]*session[^<]*id[^<]*<`)
	visibleLine          = regexp.MustCompile(`(?i)<table class="[^"]*visible[^"]*">`)
	idColumnLine         = regexp.MustCompile(`(?i)<th class="text-center">[^<]*\bid\b[^<]*<`)
	marketplaceMatch     = regexp.MustCompile(`(?i)>[^<]*marketplace[^<]*<`)
	sessionAndOrSpotLine = regexp.MustCompile(`<!-- react-text: [0-9]+ -->[1-9][0-9]*`)
	sessionAndOrSpotExpr = regexp.MustCompile(`[1-9][0-9]*$`)
	fullyPickedLine      = regexp.MustCompile(`(?i)<!-- react-text: [0-9]+ -->fully.?picked\b`)
	orderNumberLine      = regexp.MustCompile(`[0-9-]+</figcaption>`)
	orderNumberExpr      = regexp.MustCompile(`[0-9-]+`)
)

func sessionStart(html string) (s shared.Session, err error) {
	if len(html) < 1 {
		err = errors.New("no text to parse")
		return
	}

	indices := sessionLine.FindStringIndex(html)
	if len(indices) < 1 {
		err = errors.New("could not trim head; session div not found")
		return
	}
	html = html[indices[0]:]

	indices = visibleLine.FindStringIndex(html)
	if len(indices) < 1 {
		err = errors.New("could not trim tail; visible div not found")
		return
	}
	html = html[:indices[0]]

	idColumns := idColumnLine.FindAllString(html, -1)
	if len(idColumns) < 1 {
		err = errors.New("id columns not found")
		return
	}
	marketCol := -1
	for i, column := range idColumns {
		if marketplaceMatch.MatchString(column) {
			marketCol = i
			break
		}
	}
	if marketCol == -1 {
		err = errors.New("marketplace id column not found")
		return
	}

	sessionAndOrSpot := sessionAndOrSpotLine.FindAllString(html, 1)
	if len(sessionAndOrSpot) < 1 {
		err = errors.New("session id not found")
		return
	}
	sessionID, _ := strconv.Atoi(sessionAndOrSpotExpr.FindAllString(sessionAndOrSpot[0], 1)[0])

	parts := strings.Split(html, `<tr class="text-center">`)
	if len(parts) < 2 {
		err = errors.New("no rows to parse")
		return
	}
	parts = parts[1:]

	cubbies := make([]shared.Cubby, len(parts))
	for i, part := range parts {
		if !fullyPickedLine.MatchString(part) {
			println(i+1, "isn't fully picked")
			continue
		}

		// sessionAndOrSpot := sessionAndOrSpotLine.FindAllString(part, -1)
		// if len(sessionAndOrSpot) < 1 {
		// 	err = errors.New("spot not found")
		// 	return
		// }
		// spot, _ := strconv.Atoi(sessionAndOrSpotExpr.FindAllString(sessionAndOrSpot[len(sessionAndOrSpot)-1], 1)[0])

		ids := orderNumberLine.FindAllString(part, -1)
		if len(ids) < marketCol+1 {
			err = errors.New("no order id found")
			return
		}
		orderNum := orderNumberExpr.FindAllString(ids[marketCol], 1)[0]

		cubbies[i] = shared.Cubby{OrderNumber: orderNum}
	}

	s = shared.Session{
		ID:      sessionID,
		Cubbies: cubbies,
	}

	return
}

// func (sess *session) broadcast() {
// 	sess.RLock()
// 	defer sess.RUnlock()

// 	b, _ := json.Marshal(sess.Cubbies)
// 	sess.w(b)
// }
