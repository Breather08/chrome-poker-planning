package main

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type Command struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload,omitempty"`
}

type Client struct {
	User *User
	Room *Room
	Conn *websocket.Conn
	Mu   sync.Mutex
}

func NewClient(room *Room, conn *websocket.Conn) *Client {
	return &Client{
		Room: room,
		Conn: conn,
	}
}

func (c *Client) SetMessageReceiver() {
	defer func() {
		// TODO: Reconnect availability
		c.Mu.Lock()
		defer c.Mu.Unlock()
		c.Room.RemovePlayer(c.Conn, c.User)
		c.Conn.Close()
		c.SendGameState()
		fmt.Println("Player disconnected")
	}()

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		c.HandleCommand(msg)
	}
}

func (c *Client) HandlePlayerEnter(rawData json.RawMessage) {
	var payload struct {
		ID       string `json:"id"`
		Username string `json:"username"`
	}

	if err := json.Unmarshal(rawData, &payload); err != nil {
		fmt.Println("Pick failed", string(rawData))
		return
	}

	c.User = NewUser(payload.ID, payload.Username)
	c.Room.AddPlayer(c.Conn, c.User)

	req, err := json.Marshal(SerializeMessage(MessageTypePlayerEnter, c.User.ID))
	if err != nil {
		fmt.Println("Marshal failed", err)
		return
	}

	c.Room.Broadcast(req)
}

func (c *Client) HandlePlayerExit() {
	delete(c.Room.Connections, c.Conn)
}

func (c *Client) HandlePickCard(rawData json.RawMessage) {
	var payload struct {
		CardId string `json:"cardId"`
	}

	if err := json.Unmarshal(rawData, &payload); err != nil {
		fmt.Println("Pick failed", string(rawData))
		return
	}

	c.Room.GameState.PickCard(c.User, payload.CardId)
	c.SendGameState()
}

func (c *Client) HandleCardsVisibility(isRevealed bool) {
	c.Room.GameState.Revealed = isRevealed
	c.SendGameState()
}

func (c *Client) HandleReset() {
	c.Room.GameState.Reset()
	c.SendGameState()
}

func (c *Client) SendGameState() {
	msg, err := json.Marshal(SerializeMessage(MessageTypeGameState, c.Room.GameState.GetSerialized()))
	if err != nil {
		fmt.Println("Error marshalling players list", err)
		return
	}

	c.Room.Broadcast(msg)
}

func (c *Client) HandleCommand(message []byte) {
	var cmd Command
	if err := json.Unmarshal(message, &cmd); err != nil {
		fmt.Println("Invalid command format:", err)
		return
	}

	switch cmd.Type {
	case "join":
		c.HandlePlayerEnter(cmd.Payload)
	case "pick":
		c.HandlePickCard(cmd.Payload)
	case "show":
		c.HandleCardsVisibility(true)
	case "reset":
		c.HandleReset()
	case "get_state":
		c.SendGameState()
	default:
		fmt.Println("Unknown command type:", cmd.Type)
	}
}
