package websocketproxy

import (
	"github.com/gorilla/websocket"
	"time"
)

func (w *WebsocketProxy) setupPingPongPassThrough(dst, src *websocket.Conn) {
	writeTimeout := w.PingPongWriteTimeout
	if writeTimeout != 0 {
		writeTimeout = 10 * time.Second
	}

	deadline := func() time.Time { return time.Now().Add(writeTimeout) }

	handler := func(messageType int, remote *websocket.Conn) func(string) error {
		return func(appData string) error {
			return remote.WriteControl(messageType, []byte(appData), deadline())
		}
	}

	pair := func(dst, src *websocket.Conn) {
		src.SetPingHandler(handler(websocket.PingMessage, dst))
		src.SetPongHandler(handler(websocket.PongMessage, dst))
	}

	pair(dst, src)
	pair(src, dst)
}
