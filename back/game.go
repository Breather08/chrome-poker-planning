package main

type GameState struct {
	Revealed bool
	Deck     []string
	Players  map[string]*User
	Votes    map[string]string
}

func NewGameState() *GameState {
	return &GameState{
		Revealed: false,
		Votes:    map[string]string{},
		Players:  map[string]*User{},
		Deck:     []string{"1", "2", "3", "5", "8", "13"},
	}
}

func (gs *GameState) PickCard(user *User, cardId string) {
	gs.Votes[user.ID] = cardId
	user.Voted = true
	user.Vote = cardId
}
