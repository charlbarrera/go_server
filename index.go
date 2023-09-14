// start a server
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	// start a server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World! from go server!! from w.Write"))
	})
	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello from /chat"))
	})

	// handle websocket connections
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// upgrade the HTTP connection to a WebSocket connection
		upgrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// allow connections from localhost
				return true
			},
		}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Error upgrading connection to WebSocket:", err)
			return
		}
		defer conn.Close()

		// handle WebSocket messages
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message from WebSocket:", err)
				return
			}
			log.Println("Received message from WebSocket:", string(message))
			conn.WriteMessage(websocket.TextMessage, []byte("Hello from server!"))

			// handle the message
			// ...
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
