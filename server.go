package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/nfnt/resize"

	"github.com/mrmiguu/sock"

	"github.com/WedgeNix/CubbyChaser-shared"
	"github.com/WedgeNix/CubbyChaser/session"

	// _ "github.com/joho/godotenv/autoload"

	"github.com/mrmiguu/Loading"
)

var (
	newSession = make(chan shared.Session)
)

func init() {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "5000"
	}
	http.HandleFunc("/createSession", createSession)
	sock.Addr = ":" + port
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

		shared.Must(os.MkdirAll("www/"+shared.PictureFolder, os.ModePerm))
		for _, ord := range sess.Cubbies {
			for _, itm := range ord.Items {
				file := "www/" + shared.PictureFolder + "/" + itm.UPC + ".jpg"
				if _, err := os.Stat(file); !os.IsNotExist(err) {
					continue
				}
				resp, err := http.Get(itm.ImageURL)
				if err != nil {
					continue
				}
				img, _, err := image.Decode(resp.Body)
				resp.Body.Close()
				if err != nil {
					continue
				}
				f, err := os.Create(file)
				shared.Must(err)
				jpeg.Encode(f, resize.Resize(120, 0, img, resize.Bilinear), nil)
				f.Close()
			}
		}
		done <- false

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

				queuel.Lock()
				delete(queue, sess.ID)
				b := shared.Queue2bytes(queue)
				queuel.Unlock()

				Queue <- b
			}(sess)
		}
	}()

	for range Join {
		println(`[[A new user joined]]`)
		go func() {
			queuel.RLock()
			b := shared.Queue2bytes(queue)
			queuel.RUnlock()

			Queue <- b
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

	db := map[int]bool{}

	kill, Kill := make(chan bool), sock.Rbool(SOCKSession)
	UID := sock.Rint(SOCKSession)
	Found := sock.Wbool(SOCKSession)

	for {
		var uid int
		select {
		case <-Kill:
			for i := 0; i < len(db); i++ {
				kill <- true
			}
			return
		case uid = <-UID:
			var ids []int
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

		go assistUser(full, uid, &sess, Ords, kill)

		db[uid] = true
	}
}

func assistUser(full shared.Session, uid int, sess *iSession, Ords []chan<- []byte, kill <-chan bool) {
	done := load.New(`[[Assisting a new user]]`)

	SOCKSessionUser := shared.SOCKSessionUser(full.ID, uid)
	defer sock.Close(SOCKSessionUser)

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
		case <-kill:
			Bail <- true
			return
		case <-Leave:
			return
		case upc = <-UPC:
		}

		if len(upc) == 0 {
			Spot <- -1
		}

		sess.Lock()
		for spot, ord := range full.Cubbies {
			for _, orig := range ord.Items {
				if orig.UPC != upc && orig.SKU != upc {
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
