package main

type MessageType string

type Message[T any] struct {
	MessageType MessageType `json:"type"`
	MessageBody T           `json:"body"`
}

const (
	MessageTypePlayerEnter MessageType = "player_enter"
	MessageTypePickCard                = "pick_card"
	MessageTypeGameState               = "game_state"
	MessageTypePlayersList             = "players_list"
)

func SerializeMessage[T any](msgType MessageType, msgBody T) Message[T] {
	return Message[T]{
		MessageType: msgType,
		MessageBody: msgBody,
	}
}
