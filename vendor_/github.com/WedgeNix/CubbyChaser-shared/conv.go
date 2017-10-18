package shared

import (
	"encoding/json"
)

func Queue2bytes(q Queue) []byte {
	b, _ := json.Marshal(q)
	return b
}

func Bytes2queue(b []byte) Queue {
	q := Queue{}
	json.Unmarshal(b, &q)
	return q
}

func Session2bytes(s Session) []byte {
	b, _ := json.Marshal(s)
	return b
}

func Bytes2session(b []byte) Session {
	var s Session
	json.Unmarshal(b, &s)
	return s
}

func Order2bytes(o Order) []byte {
	b, _ := json.Marshal(o)
	return b
}

func Bytes2order(b []byte) Order {
	var o Order
	json.Unmarshal(b, &o)
	return o
}
