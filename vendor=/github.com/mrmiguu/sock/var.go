package sock

import (
	"sync"

	"github.com/gopherjs/gopherjs/js"
	"github.com/gorilla/websocket"
)

var (
	v = []byte(V)

	start sync.Once

	ws     *js.Object
	connl  sync.RWMutex
	conns  = map[*websocket.Conn]bool{}
	reboot sync.RWMutex
	wsync  sync.Mutex

	wbooll    sync.Mutex
	wboolm    = make(map[string][]interface{})
	rbooll    sync.RWMutex
	rboolm    = make(map[string][]rbool)
	wstringl  sync.Mutex
	wstringm  = make(map[string][]interface{})
	rstringl  sync.RWMutex
	rstringm  = make(map[string][]rstring)
	wintl     sync.Mutex
	wintm     = make(map[string][]interface{})
	rintl     sync.RWMutex
	rintm     = make(map[string][]rint)
	wint8l    sync.Mutex
	wint8m    = make(map[string][]interface{})
	rint8l    sync.RWMutex
	rint8m    = make(map[string][]rint8)
	wint16l   sync.Mutex
	wint16m   = make(map[string][]interface{})
	rint16l   sync.RWMutex
	rint16m   = make(map[string][]rint16)
	wint32l   sync.Mutex
	wint32m   = make(map[string][]interface{})
	rint32l   sync.RWMutex
	rint32m   = make(map[string][]rint32)
	wint64l   sync.Mutex
	wint64m   = make(map[string][]interface{})
	rint64l   sync.RWMutex
	rint64m   = make(map[string][]rint64)
	wuintl    sync.Mutex
	wuintm    = make(map[string][]interface{})
	ruintl    sync.RWMutex
	ruintm    = make(map[string][]ruint)
	wuint8l   sync.Mutex
	wuint8m   = make(map[string][]interface{})
	ruint8l   sync.RWMutex
	ruint8m   = make(map[string][]ruint8)
	wuint16l  sync.Mutex
	wuint16m  = make(map[string][]interface{})
	ruint16l  sync.RWMutex
	ruint16m  = make(map[string][]ruint16)
	wuint32l  sync.Mutex
	wuint32m  = make(map[string][]interface{})
	ruint32l  sync.RWMutex
	ruint32m  = make(map[string][]ruint32)
	wuint64l  sync.Mutex
	wuint64m  = make(map[string][]interface{})
	ruint64l  sync.RWMutex
	ruint64m  = make(map[string][]ruint64)
	wbytel    sync.Mutex
	wbytem    = make(map[string][]interface{})
	rbytel    sync.RWMutex
	rbytem    = make(map[string][]rbyte)
	wbytesl   sync.Mutex
	wbytesm   = make(map[string][]interface{})
	rbytesl   sync.RWMutex
	rbytesm   = make(map[string][]rbytes)
	wrunel    sync.Mutex
	wrunem    = make(map[string][]interface{})
	rrunel    sync.RWMutex
	rrunem    = make(map[string][]rrune)
	wfloat32l sync.Mutex
	wfloat32m = make(map[string][]interface{})
	rfloat32l sync.RWMutex
	rfloat32m = make(map[string][]rfloat32)
	wfloat64l sync.Mutex
	wfloat64m = make(map[string][]interface{})
	rfloat64l sync.RWMutex
	rfloat64m = make(map[string][]rfloat64)
	werrorl   sync.Mutex
	werrorm   = make(map[string][]interface{})
	rerrorl   sync.RWMutex
	rerrorm   = make(map[string][]rerror)
)

func init() {
	reboot.Lock()
}
