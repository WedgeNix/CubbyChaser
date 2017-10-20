package load

import "strings"
import "time"

func New(proc string, d ...time.Duration) (done chan<- bool) {
	c := make(chan bool)
	done = c
	println(proc + "...")
	spaces := strings.Repeat(" ", len(proc))
	go func() {
		var t time.Ticker
		if len(d) > 1 {
			panic("too many arguments")
		}
		if len(d) > 0 {
			t = *time.NewTicker(d[0])
			go func() {
				for range t.C {
					done <- false
				}
			}()
		}
		for is := range c {
			if is {
				t.Stop()
				close(c)
				println(spaces + "!!!")
				return
			}
			println(spaces + "...")
		}
	}()
	return
}

func Is(done chan<- bool) {
	done <- true
}
