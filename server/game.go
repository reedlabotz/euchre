package euchre

type Card int
type Deck [24]Card
type Hand []Card
type Action int

type Player struct {
	Name string
	Id uuid
	Hand Hand
}

type Team struct {
	Players [2]Player
	Score int
	HandsWon int
}

type Game struct {
	Teams [2]Team
	Flip Card
	Kitty [3]Card
	Dealer Player
	CurrentPlayer Player
	Table []Move
}

type Move struct {
	Player Player
	Action Action
	Card Card
}

type ClientMessageType int

const (
	HELLO ClientMessageType = 0
	MOVE ClientMessageType = 1
)

type ClientMessage struct {
	Type ClientMessageType
	Move Move
	GameHash string
}

type ServerMessageType int

const (
	GAME_REFRESH ServerMessageType = 0
	MOVE ServerMessageType = 1
)

type ServerMessage struct {
	Type ServerMessageType
	Move Move
	Game Game
	GameHash string
}