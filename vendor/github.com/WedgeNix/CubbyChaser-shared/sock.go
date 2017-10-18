package shared

import "strconv"

func SOCKSession(id int) string {
	return "session#" + strconv.Itoa(id)
}

func SOCKSessionUser(id int, uid int) string {
	return "session#" + strconv.Itoa(id) + "@" + strconv.Itoa(uid)
}

func SOCKSessionCubby(id int, spot int) string {
	return "session#" + strconv.Itoa(id) + "|cubby#" + strconv.Itoa(spot)
}
