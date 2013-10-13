package main

type Card int
type Hand []Card
type Action int

type Player struct {
	Name string
//	Id uuid
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

