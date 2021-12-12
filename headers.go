package websocketproxy

import (
	"net/http"
	"strings"
)

func copySecWebsocketHeaders(dst, src http.Header) {
	for key, values := range src {
		if !strings.HasPrefix(strings.ToLower(key), "sec-websocket-") {
			continue
		}

		for _, value := range values {
			dst.Add(key, value)
		}
	}
}
