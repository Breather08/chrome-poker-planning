package main

import "fmt"

type GameState struct {
	Revealed bool             `json:"revealed"`
	Deck     []string         `json:"-"`
	Players  map[string]*User `json:"players"`
}

type GameStateSerialized struct {
	Revealed bool   `json:"revealed"`
	Players  []User `json:"players"`
}

func NewGameState() *GameState {
	return &GameState{
		Revealed: false,
		Players:  map[string]*User{},
		Deck:     []string{"1", "2", "3", "5", "8", "13"},
	}
}

func (gs *GameState) PickCard(user *User, cardId string) {
	user.Voted = true
	user.Vote = cardId
}

func (gs *GameState) Reset() {
	gs.Revealed = false
	for _, player := range gs.Players {
		player.Vote = ""
		player.Voted = false
	}
}

func (gs *GameState) GetSerialized() GameStateSerialized {
	fmt.Println(gs.Players)
	players := make([]User, 0, len(gs.Players))
	for _, val := range gs.Players {
		players = append(players, *val)
	}

	return GameStateSerialized{
		Revealed: gs.Revealed,
		Players:  players,
	}
}
