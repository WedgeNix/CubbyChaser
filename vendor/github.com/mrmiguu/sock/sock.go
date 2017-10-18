package sock

import "github.com/gopherjs/gopherjs/js"

var (
	Addr = "localhost:80"

	IsClient = js.Global != nil

	Root = "www"

	API = "/317d37b0edc7bd7cbd25d97f53a16ce5"
)

func Close(key string) {
	wbooll.Lock()
	delete(wboolm, key)
	wbooll.Unlock()

	rbooll.Lock()
	delete(rboolm, key)
	rbooll.Unlock()

	wstringl.Lock()
	delete(wstringm, key)
	wstringl.Unlock()

	rstringl.Lock()
	delete(rstringm, key)
	rstringl.Unlock()

	wintl.Lock()
	delete(wintm, key)
	wintl.Unlock()

	rintl.Lock()
	delete(rintm, key)
	rintl.Unlock()

	wint8l.Lock()
	delete(wint8m, key)
	wint8l.Unlock()

	rint8l.Lock()
	delete(rint8m, key)
	rint8l.Unlock()

	wint16l.Lock()
	delete(wint16m, key)
	wint16l.Unlock()

	rint16l.Lock()
	delete(rint16m, key)
	rint16l.Unlock()

	wint32l.Lock()
	delete(wint32m, key)
	wint32l.Unlock()

	rint32l.Lock()
	delete(rint32m, key)
	rint32l.Unlock()

	wint64l.Lock()
	delete(wint64m, key)
	wint64l.Unlock()

	rint64l.Lock()
	delete(rint64m, key)
	rint64l.Unlock()

	wuintl.Lock()
	delete(wuintm, key)
	wuintl.Unlock()

	ruintl.Lock()
	delete(ruintm, key)
	ruintl.Unlock()

	wuint8l.Lock()
	delete(wuint8m, key)
	wuint8l.Unlock()

	ruint8l.Lock()
	delete(ruint8m, key)
	ruint8l.Unlock()

	wuint16l.Lock()
	delete(wuint16m, key)
	wuint16l.Unlock()

	ruint16l.Lock()
	delete(ruint16m, key)
	ruint16l.Unlock()

	wuint32l.Lock()
	delete(wuint32m, key)
	wuint32l.Unlock()

	ruint32l.Lock()
	delete(ruint32m, key)
	ruint32l.Unlock()

	wuint64l.Lock()
	delete(wuint64m, key)
	wuint64l.Unlock()

	ruint64l.Lock()
	delete(ruint64m, key)
	ruint64l.Unlock()

	wbytel.Lock()
	delete(wbytem, key)
	wbytel.Unlock()

	rbytel.Lock()
	delete(rbytem, key)
	rbytel.Unlock()

	wbytesl.Lock()
	delete(wbytesm, key)
	wbytesl.Unlock()

	rbytesl.Lock()
	delete(rbytesm, key)
	rbytesl.Unlock()

	wrunel.Lock()
	delete(wrunem, key)
	wrunel.Unlock()

	rrunel.Lock()
	delete(rrunem, key)
	rrunel.Unlock()

	wfloat32l.Lock()
	delete(wfloat32m, key)
	wfloat32l.Unlock()

	rfloat32l.Lock()
	delete(rfloat32m, key)
	rfloat32l.Unlock()

	wfloat64l.Lock()
	delete(wfloat64m, key)
	wfloat64l.Unlock()

	rfloat64l.Lock()
	delete(rfloat64m, key)
	rfloat64l.Unlock()

	werrorl.Lock()
	delete(werrorm, key)
	werrorl.Unlock()

	rerrorl.Lock()
	delete(rerrorm, key)
	rerrorl.Unlock()
}

func Wbool(key ...string) chan<- bool {
	k := getKey(key...)

	start.Do(run)

	w := make(chan bool)

	wbooll.Lock()
	idx := len(wboolm[k])
	wboolm[k] = append(wboolm[k], nil)
	wbooll.Unlock()

	go func() {
		for b := range w {
			write(Tbool, k, idx, bool2bytes(b))
		}
	}()

	return w
}

func Rbool(key ...string) <-chan bool {
	k := getKey(key...)

	start.Do(run)

	r := make(chan bool)

	rbooll.Lock()
	var B rbool
	B.key = k
	B.idx = len(rboolm[k])
	B.r = r
	rboolm[k] = append(rboolm[k], B)
	rbooll.Unlock()

	return r
}

func Wstring(key ...string) chan<- string {
	k := getKey(key...)

	start.Do(run)

	w := make(chan string)

	wstringl.Lock()
	idx := len(wstringm[k])
	wstringm[k] = append(wstringm[k], nil)
	wstringl.Unlock()

	go func() {
		for s := range w {
			write(Tstring, k, idx, []byte(s))
		}
	}()

	return w
}

func Rstring(key ...string) <-chan string {
	k := getKey(key...)

	start.Do(run)

	r := make(chan string)

	rstringl.Lock()
	var S rstring
	S.key = k
	S.idx = len(rstringm[k])
	S.r = r
	rstringm[k] = append(rstringm[k], S)
	rstringl.Unlock()

	return r
}

func Wint(key ...string) chan<- int {
	k := getKey(key...)

	start.Do(run)

	w := make(chan int)

	wintl.Lock()
	idx := len(wintm[k])
	wintm[k] = append(wintm[k], nil)
	wintl.Unlock()

	go func() {
		for i := range w {
			write(Tint, k, idx, int2bytes(i))
		}
	}()

	return w
}

func Rint(key ...string) <-chan int {
	k := getKey(key...)

	start.Do(run)

	r := make(chan int)

	rintl.Lock()
	var I rint
	I.key = k
	I.idx = len(rintm[k])
	I.r = r
	rintm[k] = append(rintm[k], I)
	rintl.Unlock()

	return r
}

func Wint8(key ...string) chan<- int8 {
	k := getKey(key...)

	start.Do(run)

	w := make(chan int8)

	wint8l.Lock()
	idx := len(wint8m[k])
	wint8m[k] = append(wint8m[k], nil)
	wint8l.Unlock()

	go func() {
		for i := range w {
			write(Tint8, k, idx, int82bytes(i))
		}
	}()

	return w
}

func Rint8(key ...string) <-chan int8 {
	k := getKey(key...)

	start.Do(run)

	r := make(chan int8)

	rint8l.Lock()
	var I rint8
	I.key = k
	I.idx = len(rint8m[k])
	I.r = r
	rint8m[k] = append(rint8m[k], I)
	rint8l.Unlock()

	return r
}

func Wint16(key ...string) chan<- int16 {
	k := getKey(key...)

	start.Do(run)

	w := make(chan int16)

	wint16l.Lock()
	idx := len(wint16m[k])
	wint16m[k] = append(wint16m[k], nil)
	wint16l.Unlock()

	go func() {
		for i := range w {
			write(Tint16, k, idx, int162bytes(i))
		}
	}()

	return w
}

func Rint16(key ...string) <-chan int16 {
	k := getKey(key...)

	start.Do(run)

	r := make(chan int16)

	rint16l.Lock()
	var I rint16
	I.key = k
	I.idx = len(rint16m[k])
	I.r = r
	rint16m[k] = append(rint16m[k], I)
	rint16l.Unlock()

	return r
}

func Wint32(key ...string) chan<- int32 {
	k := getKey(key...)

	start.Do(run)

	w := make(chan int32)

	wint32l.Lock()
	idx := len(wint32m[k])
	wint32m[k] = append(wint32m[k], nil)
	wint32l.Unlock()

	go func() {
		for i := range w {
			write(Tint32, k, idx, int322bytes(i))
		}
	}()

	return w
}

func Rint32(key ...string) <-chan int32 {
	k := getKey(key...)

	start.Do(run)

	r := make(chan int32)

	rint32l.Lock()
	var I rint32
	I.key = k
	I.idx = len(rint32m[k])
	I.r = r
	rint32m[k] = append(rint32m[k], I)
	rint32l.Unlock()

	return r
}

func Wint64(key ...string) chan<- int64 {
	k := getKey(key...)

	start.Do(run)

	w := make(chan int64)

	wint64l.Lock()
	idx := len(wint64m[k])
	wint64m[k] = append(wint64m[k], nil)
	wint64l.Unlock()

	go func() {
		for i := range w {
			write(Tint64, k, idx, int642bytes(i))
		}
	}()

	return w
}

func Rint64(key ...string) <-chan int64 {
	k := getKey(key...)

	start.Do(run)

	r := make(chan int64)

	rint64l.Lock()
	var I rint64
	I.key = k
	I.idx = len(rint64m[k])
	I.r = r
	rint64m[k] = append(rint64m[k], I)
	rint64l.Unlock()

	return r
}

func Wuint(key ...string) chan<- uint {
	k := getKey(key...)

	start.Do(run)

	w := make(chan uint)

	wuintl.Lock()
	idx := len(wuintm[k])
	wuintm[k] = append(wuintm[k], nil)
	wuintl.Unlock()

	go func() {
		for u := range w {
			write(Tuint, k, idx, uint2bytes(u))
		}
	}()

	return w
}

func Ruint(key ...string) <-chan uint {
	k := getKey(key...)

	start.Do(run)

	r := make(chan uint)

	ruintl.Lock()
	var U ruint
	U.key = k
	U.idx = len(ruintm[k])
	U.r = r
	ruintm[k] = append(ruintm[k], U)
	ruintl.Unlock()

	return r
}

func Wuint8(key ...string) chan<- uint8 {
	k := getKey(key...)

	start.Do(run)

	w := make(chan uint8)

	wuint8l.Lock()
	idx := len(wuint8m[k])
	wuint8m[k] = append(wuint8m[k], nil)
	wuint8l.Unlock()

	go func() {
		for u := range w {
			write(Tuint8, k, idx, uint82bytes(u))
		}
	}()

	return w
}

func Ruint8(key ...string) <-chan uint8 {
	k := getKey(key...)

	start.Do(run)

	r := make(chan uint8)

	ruint8l.Lock()
	var U ruint8
	U.key = k
	U.idx = len(ruint8m[k])
	U.r = r
	ruint8m[k] = append(ruint8m[k], U)
	ruint8l.Unlock()

	return r
}

func Wuint16(key ...string) chan<- uint16 {
	k := getKey(key...)

	start.Do(run)

	w := make(chan uint16)

	wuint16l.Lock()
	idx := len(wuint16m[k])
	wuint16m[k] = append(wuint16m[k], nil)
	wuint16l.Unlock()

	go func() {
		for u := range w {
			write(Tuint16, k, idx, uint162bytes(u))
		}
	}()

	return w
}

func Ruint16(key ...string) <-chan uint16 {
	k := getKey(key...)

	start.Do(run)

	r := make(chan uint16)

	ruint16l.Lock()
	var U ruint16
	U.key = k
	U.idx = len(ruint16m[k])
	U.r = r
	ruint16m[k] = append(ruint16m[k], U)
	ruint16l.Unlock()

	return r
}

func Wuint32(key ...string) chan<- uint32 {
	k := getKey(key...)

	start.Do(run)

	w := make(chan uint32)

	wuint32l.Lock()
	idx := len(wuint32m[k])
	wuint32m[k] = append(wuint32m[k], nil)
	wuint32l.Unlock()

	go func() {
		for u := range w {
			write(Tuint32, k, idx, uint322bytes(u))
		}
	}()

	return w
}

func Ruint32(key ...string) <-chan uint32 {
	k := getKey(key...)

	start.Do(run)

	r := make(chan uint32)

	ruint32l.Lock()
	var U ruint32
	U.key = k
	U.idx = len(ruint32m[k])
	U.r = r
	ruint32m[k] = append(ruint32m[k], U)
	ruint32l.Unlock()

	return r
}

func Wuint64(key ...string) chan<- uint64 {
	k := getKey(key...)

	start.Do(run)

	w := make(chan uint64)

	wuint64l.Lock()
	idx := len(wuint64m[k])
	wuint64m[k] = append(wuint64m[k], nil)
	wuint64l.Unlock()

	go func() {
		for u := range w {
			write(Tuint64, k, idx, uint642bytes(u))
		}
	}()

	return w
}

func Ruint64(key ...string) <-chan uint64 {
	k := getKey(key...)

	start.Do(run)

	r := make(chan uint64)

	ruint64l.Lock()
	var U ruint64
	U.key = k
	U.idx = len(ruint64m[k])
	U.r = r
	ruint64m[k] = append(ruint64m[k], U)
	ruint64l.Unlock()

	return r
}

func Wbyte(key ...string) chan<- byte {
	k := getKey(key...)

	start.Do(run)

	w := make(chan byte)

	wbytel.Lock()
	idx := len(wbytem[k])
	wbytem[k] = append(wbytem[k], nil)
	wbytel.Unlock()

	go func() {
		for b := range w {
			write(Tbyte, k, idx, []byte{b})
		}
	}()

	return w
}

func Rbyte(key ...string) <-chan byte {
	k := getKey(key...)

	start.Do(run)

	r := make(chan byte)

	rbytel.Lock()
	var B rbyte
	B.key = k
	B.idx = len(rbytem[k])
	B.r = r
	rbytem[k] = append(rbytem[k], B)
	rbytel.Unlock()

	return r
}

func Wbytes(key ...string) chan<- []byte {
	k := getKey(key...)

	start.Do(run)

	w := make(chan []byte)

	wbytesl.Lock()
	idx := len(wbytesm[k])
	wbytesm[k] = append(wbytesm[k], nil)
	wbytesl.Unlock()

	go func() {
		for b := range w {
			write(Tbytes, k, idx, b)
		}
	}()

	return w
}

func Rbytes(key ...string) <-chan []byte {
	k := getKey(key...)

	start.Do(run)

	r := make(chan []byte)

	rbytesl.Lock()
	var B rbytes
	B.key = k
	B.idx = len(rbytesm[k])
	B.r = r
	rbytesm[k] = append(rbytesm[k], B)
	rbytesl.Unlock()

	return r
}

func Wrune(key ...string) chan<- rune {
	k := getKey(key...)

	start.Do(run)

	w := make(chan rune)

	wrunel.Lock()
	idx := len(wrunem[k])
	wrunem[k] = append(wrunem[k], nil)
	wrunel.Unlock()

	go func() {
		for r := range w {
			write(Trune, k, idx, rune2bytes(r))
		}
	}()

	return w
}

func Rrune(key ...string) <-chan rune {
	k := getKey(key...)

	start.Do(run)

	r := make(chan rune)

	rrunel.Lock()
	var R rrune
	R.key = k
	R.idx = len(rrunem[k])
	R.r = r
	rrunem[k] = append(rrunem[k], R)
	rrunel.Unlock()

	return r
}

func Wfloat32(key ...string) chan<- float32 {
	k := getKey(key...)

	start.Do(run)

	w := make(chan float32)

	wfloat32l.Lock()
	idx := len(wfloat32m[k])
	wfloat32m[k] = append(wfloat32m[k], nil)
	wfloat32l.Unlock()

	go func() {
		for f := range w {
			write(Tfloat32, k, idx, float322bytes(f))
		}
	}()

	return w
}

func Rfloat32(key ...string) <-chan float32 {
	k := getKey(key...)

	start.Do(run)

	r := make(chan float32)

	rfloat32l.Lock()
	var F rfloat32
	F.key = k
	F.idx = len(rfloat32m[k])
	F.r = r
	rfloat32m[k] = append(rfloat32m[k], F)
	rfloat32l.Unlock()

	return r
}

func Wfloat64(key ...string) chan<- float64 {
	k := getKey(key...)

	start.Do(run)

	w := make(chan float64)

	wfloat64l.Lock()
	idx := len(wfloat64m[k])
	wfloat64m[k] = append(wfloat64m[k], nil)
	wfloat64l.Unlock()

	go func() {
		for f := range w {
			write(Tfloat64, k, idx, float642bytes(f))
		}
	}()

	return w
}

func Rfloat64(key ...string) <-chan float64 {
	k := getKey(key...)

	start.Do(run)

	r := make(chan float64)

	rfloat64l.Lock()
	var F rfloat64
	F.key = k
	F.idx = len(rfloat64m[k])
	F.r = r
	rfloat64m[k] = append(rfloat64m[k], F)
	rfloat64l.Unlock()

	return r
}

func Werror(key ...string) chan<- error {
	k := getKey(key...)

	start.Do(run)

	w := make(chan error)

	werrorl.Lock()
	idx := len(werrorm[k])
	werrorm[k] = append(werrorm[k], nil)
	werrorl.Unlock()

	go func() {
		for e := range w {
			write(Terror, k, idx, error2bytes(e))
		}
	}()

	return w
}

func Rerror(key ...string) <-chan error {
	k := getKey(key...)

	start.Do(run)

	r := make(chan error)

	rerrorl.Lock()
	var E rerror
	E.key = k
	E.idx = len(rerrorm[k])
	E.r = r
	rerrorm[k] = append(rerrorm[k], E)
	rerrorl.Unlock()

	return r
}

func getKey(key ...string) string {
	if len(key) == 0 {
		return ""
	}
	if len(key) > 1 {
		panic("too many arguments")
	}
	if len(key[0]) == 0 {
		panic("empty key")
	}
	return key[0]
}
