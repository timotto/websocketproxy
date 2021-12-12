package websocketproxy

import (
	"github.com/gorilla/websocket"
	"time"
)

func (w *WebsocketProxy) setupPingPongPassThrough(dst, src *websocket.Conn) {
	deadline := deadlineFunction(w.PingPongWriteTimeout)

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

func deadlineFunction(timeout time.Duration) func() time.Time {
	if timeout == 0 {
		return func() time.Time {
			return time.Time{}
		}
	}

	return func() time.Time {
		return time.Now().Add(timeout)
	}
}
