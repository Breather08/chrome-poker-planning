package main

import (
	"encoding/json"
	"fmt"

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
}

func NewClient(room *Room, conn *websocket.Conn) *Client {
	return &Client{
		Room: room,
		Conn: conn,
	}
}

func (c *Client) SetMessageReceiver() {
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

	c.SendPlayersState()
}

func (c *Client) HandleCardsVisibility() {
	c.Room.GameState.Revealed = !c.Room.GameState.Revealed
	c.SendPlayersState()
}

func (c *Client) SendPlayersState() {
	msg, err := json.Marshal(SerializeMessage(MessageTypePlayersList, c.Room.GameState.Players))
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
		c.HandleCardsVisibility()
	case "players":
		c.SendPlayersState()
	default:
		fmt.Println("Unknown command type:", cmd.Type)
	}
}
