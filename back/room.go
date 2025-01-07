package main

import (
	"time"

	"github.com/gorilla/websocket"
)

type Room struct {
	Id          string
	Connections map[*websocket.Conn]*User
	Clients     map[string]*Client
	GameState   *GameState
	CreatedAt   time.Time
}

func NewRoom(id string) *Room {
	return &Room{
		Id:          id,
		Connections: map[*websocket.Conn]*User{},
		GameState:   NewGameState(),
		CreatedAt:   time.Now(),
	}
}

func (r *Room) AddPlayer(conn *websocket.Conn, user *User) {
	r.Connections[conn] = user
	r.GameState.Players[user.ID] = user
}

func (r *Room) RemovePlayer(conn *websocket.Conn, user *User) {
	delete(r.Connections, conn)
	delete(r.GameState.Players, user.ID)
}

func (r *Room) Broadcast(message []byte) {
	for conn := range r.Connections {
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			conn.Close()
			delete(r.Connections, conn)
		}
	}
}
