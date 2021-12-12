package websocketproxy

import (
	"net/http"
	"strings"
)

func copySecWebsocketHeaders(dst, src http.Header) {
	for key, values := range src {
		lKey := strings.ToLower(key)
		if !strings.HasPrefix(lKey, "sec-websocket-") {
			continue
		}

		switch lKey {
		case "sec-websocket-key":
			continue
		case "sec-websocket-version":
			continue
		}

		for _, value := range values {
			dst.Add(key, value)
		}
	}
}
