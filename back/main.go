package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	rooms    = make(map[string]*Room) // TODO: User Redis
	upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
)

func handleWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	roomId := r.URL.Query().Get("room_id")
	if roomId == "" {
		return
	}

	room, exists := rooms[roomId]
	if !exists {
		room = NewRoom(roomId)
		rooms[roomId] = room
	}

	cl := NewClient(room, conn)
	cl.SetMessageReceiver()
}

func main() {
	http.HandleFunc("/ws", handleWebsocket)

	fmt.Println("WebSocket server started at ws://localhost:8080/ws")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
