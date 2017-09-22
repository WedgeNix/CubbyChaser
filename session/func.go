package session

import (
	"errors"
	"strconv"
	"strings"

	"github.com/WedgeNix/CubbyChaser-shared"
	"github.com/WedgeNix/CubbyChaser/ship"
)

func ParseIDAndOrderNumbers(html string) (id int, ordNums []string, err error) {
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

	sessOrSpots := sessionIDLine.FindAllString(html, 1)
	if len(sessOrSpots) < 1 {
		err = errors.New("session id not found")
		return
	}
	id, _ = strconv.Atoi(sessionIDExpr.FindAllString(sessOrSpots[0], 1)[0])

	parts := strings.Split(html, `<tr class="text-center">`)
	if len(parts) < 2 {
		err = errors.New("no rows to parse")
		return
	}
	parts = parts[1:]

	ordNums = make([]string, len(parts))
	for i, part := range parts {
		if !fullyPickedLine.MatchString(part) {
			println(i+1, "isn't fully picked")
			continue
		}

		cols := orderNumberLine.FindAllString(part, -1)
		if len(cols) < marketCol+1 {
			err = errors.New("no order id found")
			return
		}
		ordNums[i] = orderNumberExpr.FindAllString(cols[marketCol], 1)[0]
	}

	return
}

func New(id int, ordNums []string) (shared.Session, error) {
	var sess shared.Session

	sess.ID = id

	ords, err := ship.New().GetOrders(ordNums)
	if err != nil {
		return sess, err
	}
	sess.Cubbies = ords

	return sess, nil
}
