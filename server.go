package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/mrmiguu/sock"

	"github.com/WedgeNix/CubbyChaser-shared"
	"github.com/WedgeNix/CubbyChaser/session"

	// _ "github.com/joho/godotenv/autoload"

	"github.com/mrmiguu/Loading"
)

type extJSON struct {
	ID   string
	Ords []string
}

var (
	newSession = make(chan shared.Session)
)

func init() {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "5000"
	}
	http.HandleFunc("/createSession", createSession)

	http.HandleFunc("/client/client.js.map", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "client/client.js.map")
	})

	gzipClient()
	http.HandleFunc("/client.js.gz", getClient)

	sock.Addr = ":" + port
}

func gzipClient() {
	f, err := os.Create("client.js.gz")
	if err != nil {
		panic(err)
	}
	zw := gzip.NewWriter(f)
	defer zw.Close()

	b, err := ioutil.ReadFile("client/client.js")
	if err != nil {
		panic(err)
	}
	_, err = zw.Write(b)
	if err != nil {
		panic(err)
	}
}

func getClient(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("client.js.gz")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/javascript")
	w.Header().Add("Content-Encoding", "gzip")
	w.Write(data)
}

func createSession(w http.ResponseWriter, r *http.Request) {
	done := load.New("creating session")
	jRresp := extJSON{}
	err := json.NewDecoder(r.Body).Decode(&jRresp)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	done <- false

	id, err := strconv.Atoi(jRresp.ID)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	ordNums := jRresp.Ords
	done <- false

	http.Error(w, http.StatusText(http.StatusOK), http.StatusOK)
	done <- false

	go func() {
		Cap := len(ordNums)
		if Cap > 20 {
			Cap = 20
		}
		q := ordNums[:Cap]
		println("id", id)
		if id == 207 {
			q = ordNums[20:]
		}
		sess, err := session.New(id, q)
		if err != nil {
			println(err.Error())
			return
		}
		done <- false

		// shared.Must(os.MkdirAll("www/"+shared.PictureFolder, os.ModePerm))
		// for _, ord := range sess.Cubbies {
		// 	for _, itm := range ord.Items {
		// 		file := "www/" + shared.PictureFolder + "/" + itm.UPC + ".jpg"
		// 		if _, err := os.Stat(file); !os.IsNotExist(err) {
		// 			continue
		// 		}
		// 		resp, err := http.Get(itm.ImageURL)
		// 		if err != nil {
		// 			continue
		// 		}
		// 		img, _, err := image.Decode(resp.Body)
		// 		resp.Body.Close()
		// 		if err != nil {
		// 			continue
		// 		}
		// 		f, err := os.Create(file)
		// 		shared.Must(err)
		// 		jpeg.Encode(f, resize.Resize(120, 0, img, resize.Bilinear), nil)
		// 		f.Close()
		// 	}
		// }
		// done <- false

		println("ItemCount", sess.ItemCount())

		newSession <- sess
		done <- true
	}()
}

func main() {
	var queuel sync.RWMutex
	queue := shared.Queue{}

	Join := sock.Rbool()
	Queue := sock.Wbytes()

	go func() {
		for sess := range newSession {
			queuel.Lock()
			queue[sess.ID] = sess.ItemCount()
			b := shared.Queue2bytes(queue)
			queuel.Unlock()

			Queue <- b

			go func(sess shared.Session) {
				deliverSession(sess)

				done := load.New("deleting session #" + strconv.Itoa(sess.ID))

				queuel.Lock()
				delete(queue, sess.ID)
				b := shared.Queue2bytes(queue)
				queuel.Unlock()
				done <- false

				Queue <- b
				done <- true
			}(sess)
		}
	}()

	for range Join {
		println(`[[A new user joined]]`)
		go func() {
			done := load.New("sending queue to new user")
			queuel.RLock()
			b := shared.Queue2bytes(queue)
			queuel.RUnlock()
			done <- false

			Queue <- b
			done <- true
		}()
	}
}

type iSession struct {
	sync.Mutex
	shared      shared.Session
	upcSpotHold map[string][]int
}

func deliverSession(full shared.Session) {
	SOCKSession := shared.SOCKSession(full.ID)
	defer sock.Close(SOCKSession)

	println(full.String())

	sess := iSession{
		shared: shared.Session{
			ID:      full.ID,
			Cubbies: make([]shared.Order, len(full.Cubbies)),
		},
		upcSpotHold: map[string][]int{},
	}

	var Ords []chan<- []byte
	for spot, orig := range full.Cubbies {
		cubs := sess.shared.Cubbies
		cubs[spot].OrderNumber = orig.OrderNumber
		cubs[spot].Items = make([]shared.Item, len(orig.Items))
		for idx := range cubs[spot].Items {
			cubs[spot].Items[idx].UPC = orig.Items[idx].UPC
			cubs[spot].Items[idx].SKU = orig.Items[idx].SKU
			cubs[spot].Items[idx].Quantity = 0
		}

		SOCKSessionCubby := shared.SOCKSessionCubby(full.ID, spot)
		defer sock.Close(SOCKSessionCubby)
		Ords = append(Ords, sock.Wbytes(SOCKSessionCubby))
	}

	var dl sync.RWMutex
	db := map[int]chan bool{}

	kill, Kill := make(chan bool), sock.Rbool(SOCKSession)
	UID := sock.Rint(SOCKSession)
	Found := sock.Wbool(SOCKSession)

	for {
		var uid int
		select {
		case <-Kill:
			done := load.New("killing")
			dl.RLock()
			done <- false
			for i := 0; i < len(db); i++ {
				kill <- true
				done <- false
			}
			dl.RUnlock()
			done <- true
			return
		case uid = <-UID:
			var ids []int
			dl.Lock()
			for userID := range db {
				ids = append(ids, userID)
			}
			fmt.Println(ids)
		}

		_, found := db[uid]
		Found <- found
		if found {
			continue
		}

		SOCKUID := strconv.Itoa(uid)
		Ping, Pong := sock.Wbool(SOCKUID), sock.Rbool(SOCKUID)
		go func() {
			defer sock.Close(SOCKUID)
			for {
				timesup := time.NewTimer(5 * time.Second)
				select {
				case <-timesup.C:
					timesup.Stop()
					dl.RLock()
					db[uid] <- true
					dl.RUnlock()
					return
				case <-Pong:
					time.Sleep(2 * time.Second)
					Ping <- true
				}
				timesup.Stop()
			}
		}()

		go assistUser(full, uid, &sess, Ords, kill, db, &dl)

		db[uid] = make(chan bool)
		dl.Unlock()
	}
}

func closeUser(uid int, db map[int]chan bool, dl *sync.RWMutex) {
	done := load.New(`removing user #` + strconv.Itoa(uid))
	dl.Lock()
	done <- false
	delete(db, uid)
	dl.Unlock()
	done <- true
}

func assistUser(full shared.Session, uid int, sess *iSession, Ords []chan<- []byte, kill <-chan bool, db map[int]chan bool, dl *sync.RWMutex) {
	SOCKSessionUser := shared.SOCKSessionUser(full.ID, uid)
	defer sock.Close(SOCKSessionUser)
	println(SOCKSessionUser)

	done := load.New(`adding user #` + strconv.Itoa(uid))
	defer closeUser(uid, db, dl)

	Sess := sock.Wbytes(SOCKSessionUser)
	done <- false
	Sess <- shared.Session2bytes(full)
	done <- false
	sess.Lock()
	s := sess.shared
	done <- false
	sess.Unlock()
	Sess <- shared.Session2bytes(s)
	done <- false

	UPC := sock.Rstring(SOCKSessionUser)
	done <- false
	Spot := sock.Wint(SOCKSessionUser)
	done <- false
	Put := sock.Rbool(SOCKSessionUser)
	done <- false
	Cancel := sock.Wbool(SOCKSessionUser)
	done <- false
	Bail := sock.Wbool(SOCKSessionUser)
	done <- false
	Leave := sock.Rbool(SOCKSessionUser)
	done <- true

nextUPC:
	for {
		var upc string
		select {
		case <-db[uid]:
			done := load.New("bailing #" + strconv.Itoa(uid))
			time.Sleep(2 * time.Second)
			done <- true
			return
		case <-kill:
			done := load.New("bailing #" + strconv.Itoa(uid))
			Bail <- true
			done <- false
			time.Sleep(2 * time.Second)
			done <- true
			return
		case <-Leave:
			return
		case upc = <-UPC:
		}
		upc = strings.ToUpper(upc)

		if len(upc) == 0 {
			Spot <- -1
		}

		sess.Lock()
		for spot, ord := range full.Cubbies {
			for _, orig := range ord.Items {
				if orig.UPC != upc && strings.ToUpper(orig.SKU) != upc {
					continue
				}
				cubs := sess.shared.Cubbies
				for idx, itm := range cubs[spot].Items {
					if len(itm.UPC) > 0 && itm.UPC != orig.UPC {
						continue
					} else if len(itm.SKU) > 0 && itm.SKU != orig.SKU {
						continue
					}
					if itm.Quantity == orig.Quantity {
						continue
					}
					qt := itm.Quantity + 1
					cubs[spot].Items[idx] = orig
					cubs[spot].Items[idx].Quantity = qt
					b := shared.Order2bytes(cubs[spot])
					sess.Unlock()

					Ords[spot] <- b
					Spot <- spot

					timeout := time.NewTimer(shared.Timeout)
					select {
					case <-timeout.C:
						Cancel <- true
					case put := <-Put:
						if put {
							timeout.Stop()
							continue nextUPC
						}
					}
					sess.Lock()
					cubs[spot].Items[idx] = itm
					b = shared.Order2bytes(cubs[spot])
					sess.Unlock()
					Ords[spot] <- b

					timeout.Stop()
					continue nextUPC
				}
			}
		}
		sess.Unlock()

		Spot <- -1
	}
}
