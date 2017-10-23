package main

import (
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/mrmiguu/Loading"

	"github.com/WedgeNix/CubbyChaser-shared"
	"github.com/WedgeNix/CubbyChaser/www/empty"
	"github.com/gopherjs/gopherjs/js"
	"github.com/mrmiguu/jsutil"
	"github.com/mrmiguu/sock"
)

var (
	document *js.Object
	idc      = make(chan int, 100)
	qtc      []atomic.Value
)

func init() {
	document = js.Global.Get("document")
	js.Global.Call("endMainLoader")
	sock.Addr = js.Global.Get("location").Get("host").String()
}

func populateQueue(q shared.Queue) {
	js.Global.Call("populateSess", q)
	for id := range q {
		getElementById("sessBar-"+strconv.Itoa(id)).Set("onclick", func() {
			go func() {
				done := load.New("idc <- id")
				idc <- id
				done <- true
			}()
		})
	}
}

func getElementById(id string) *js.Object {
	return document.Call("getElementById", id)
}

func main() {
	var queue atomic.Value
	var readChoice sync.Once
	choice := make(chan int, 100)

	Join := sock.Wbool()
	Join <- true

	Queue := sock.Rbytes()
	for {
		select {
		case id := <-choice:
			joinSession(id)

		case b := <-Queue:
			q := shared.Bytes2queue(b)
			queue.Store(q)
			println(q.String())
			populateQueue(q)

			go readChoice.Do(func() {
				for {
					done := load.New("id := <-idc")
					id := <-idc
					done <- true
					if _, found := queue.Load().(shared.Queue)[id]; !found {
						alert("Wrong Session", "Session chosen has expired")
						continue
					}
					js.Global.Call("showLoader")
					done = load.New(`choice <- id`)
					choice <- id
					done <- true
				}
			})
		}
	}
}

func joinSession(id int) {
	done := load.New(`joinSession`)
	defer println(`[joinSession DONE]`)

	SOCKSession := shared.SOCKSession(id)
	defer sock.Close(SOCKSession)

	Kill := sock.Wbool(SOCKSession)
	done <- false
	UID := sock.Wint(SOCKSession)
	done <- false
	Found := sock.Rbool(SOCKSession)
	done <- false

	var uid int
	for gen := true; gen; gen = <-Found {
		uid = rand.Int()
		UID <- uid
		println(`uid`, uid)
	}
	done <- true

	SOCKUID := strconv.Itoa(uid)
	Ping, Pong := sock.Wbool(SOCKUID), sock.Rbool(SOCKUID)
	defer sock.Close(SOCKUID)
	complete := make(chan bool)

	go func() {
		for {
			Ping <- true
			select {
			case <-complete:
				println("[completed pinging this session]")
				return
			case <-Pong:
			}
		}
	}()

	manuallyPopulateCubbies(id, uid, Kill)
	complete <- true
}

func clearCubz() {
	getElementById("cubbies").Set("innerHTML", empty.Cubbies)
	getElementById("sess-drop").Set("innerHTML", `<i class="material-icons">keyboard_arrow_down</i> Sessions `)
	getElementById("show-sess").Call("removeAttribute", "disabled")
	getElementById("end-sess").Call("setAttribute", "disabled", "")
}

func manuallyPopulateCubbies(id, uid int, Kill chan<- bool) {
	done := load.New(`manuallyPopulateCubbies`)
	defer println(`[manuallyPopulateCubbies DONE]`)

	SOCKSessionUser := shared.SOCKSessionUser(id, uid)
	defer sock.Close(SOCKSessionUser)
	println(SOCKSessionUser)

	Sess := sock.Rbytes(SOCKSessionUser)
	full := shared.Bytes2session(<-Sess)
	done <- false
	sess := shared.Bytes2session(<-Sess)
	done <- false
	println(sess.String())
	done <- true

	qtc = make([]atomic.Value, len(full.Cubbies))
	js.Global.Call("populateCubbies", full)
	js.Global.Call("rippleCopy")
	js.Global.Call("preloadImages", full.Cubbies)
	for spot, orig := range full.Cubbies {
		updateOrder(spot, sess.Cubbies[spot], orig)
	}

	bail := make(chan bool)
	for spot := range sess.Cubbies {
		go syncOrder(id, spot, full.Cubbies[spot], bail)
	}

	getElementById("delete-sess").Set("onclick", func() {
		Kill <- true
		js.Global.Call("closeEnd")
	})

	UPC := sock.Wstring(SOCKSessionUser)
	Spot := sock.Rint(SOCKSessionUser)
	Put := sock.Wbool(SOCKSessionUser)
	Cancel := sock.Rbool(SOCKSessionUser)
	Bail := sock.Rbool(SOCKSessionUser)
	Leave := sock.Wbool(SOCKSessionUser)

	leave := make(chan bool, 1)
	getElementById("exit-sess").Set("onclick", func() {
		Leave <- true
		leave <- true
		js.Global.Call("closeEnd")
	})

	upcc := make(chan string)
	upcSKU := getElementById("upc-sku")
	upcSKU.Set("onkeypress", jsutil.F(func(e ...*js.Object) {
		if e[0].Get("keyCode").Int() == 13 {
			upcc <- e[0].Get("target").Get("value").String()
		}
	}))
	spotc := make(chan string)
	cubby := getElementById("cubby")
	cubby.Set("onkeypress", jsutil.F(func(e ...*js.Object) {
		if e[0].Get("keyCode").Int() == 13 {
			spotc <- e[0].Get("target").Get("value").String()
		}
	}))

	for {
		var upc string
		select {
		case <-leave:
			for range sess.Cubbies {
				bail <- true
			}
			clearCubz()
			return
		case <-Bail:
			for range sess.Cubbies {
				bail <- true
			}
			clearCubz()
			return
		case upc = <-upcc:
		}

		UPC <- upc

		spot := <-Spot
		if spot == -1 {
			js.Global.Call("clryou")
			alert("Wrong UPC/SKU", "No match for <b>"+upc+"</b>; try using SKU or re-enter the UPC")
			continue
		}

		n := strconv.Itoa(spot + 1)
		D000 := "D" + strings.Repeat("0", 4-len(n)) + n

		println("spot:", D000)
		js.Global.Call("sendToCubby", shared.PictureFolder+"/"+upc+".jpg", spot)

		shortCircuit := make(chan bool)
		clickwall := getElementById("clickwall")
		clickwall.Call("removeAttribute", "hidden")
		clickwall.Set("onclick", jsutil.F(func(_ ...*js.Object) { shortCircuit <- true }))
		js.Global.Get("window").Set("onbeforeunload", jsutil.F(func(_ ...*js.Object) { shortCircuit <- true }))

	nextLoc:
		for {
			select {
			case <-shortCircuit:
				Put <- false
			case <-Cancel:
			case loc := <-spotc:
				if loc != D000 {
					js.Global.Call("clrcub")
					alert("Wrong Cubby", "<b>"+loc+"</b> is not the right cubby. Check that you put the product into the right cubby")
					continue nextLoc
				}
				select {
				case <-Cancel:
				case Put <- true:
				}
			}
			break
		}

		js.Global.Call("stopShake", spot, qtc[spot].Load().(int), full.Cubbies[spot].ItemCount())
	}
}

func syncOrder(id, spot int, orig shared.Order, bail <-chan bool) {
	SOCKSessionCubby := shared.SOCKSessionCubby(id, spot)
	defer sock.Close(SOCKSessionCubby)

	Ord := sock.Rbytes(SOCKSessionCubby)
	for {
		var b []byte
		select {
		case <-bail:
			return
		case b = <-Ord:
		}

		ord := shared.Bytes2order(b)
		println(ord.String())
		updateOrder(spot, ord, orig)
		qtc[spot].Store(ord.ItemCount())
	}
}

func updateOrder(spot int, ord, orig shared.Order) {
	js.Global.Call("orderCount", spot, ord.ItemCount(), orig.ItemCount())
}

func alert(title, message string) {
	js.Global.Call("materialAlert", title, message)
}
