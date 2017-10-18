package sock

import (
	"bytes"

	"github.com/gorilla/websocket"
)

func write(t byte, key string, idx int, body []byte) {
	pkt := bytes.Join([][]byte{[]byte{t}, []byte(key), int2bytes(idx), body}, v)

	if IsClient {
		wsync.Lock()
		ws.Call("send", pkt)
		wsync.Unlock()
		return
	}
	for {
		connl.RLock()
		cnt := len(conns)
		for conn := range conns {
			wsync.Lock()
			conn.WriteMessage(websocket.BinaryMessage, pkt)
			wsync.Unlock()
		}
		connl.RUnlock()

		if cnt > 0 {
			break
		}

		reboot.RLock()
		reboot.RUnlock()
	}
}
