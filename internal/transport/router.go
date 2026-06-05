package transport

import (
	"net/http"

	"github.com/coder/websocket"
)

func NewRouter(wsHandler *WSHandler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Accept(w, r, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		wsHandler.HandleConnection(r.Context(), conn)
		conn.Close(websocket.StatusNormalClosure, "OK")
	})

	return mux
}
