package main

import (
	"log"
	"time"
)

type Storage struct {
	ttl int                // Time to live in ms for a game
	ticker *time.Ticker    // Cleanup channel
	games map[string] Game // Make of GameIds to games
}

func NewStorage(ttl time.Duration) *Storage {
	s := new(Storage)
	s.ticker = time.NewTicker(ttl/2)
	go s.Cleanup()
	return s
}

func (s *Storage) Cleanup() {
	for {
		select {
		case <- s.ticker.C:
			log.Printf("Running cleanup")
		}
	}
}

func (s *Storage) GetGame(id string) (Game, error) {
	return s.games[id], nil
}

func (s *Storage) StartGame(id string) error {
	return nil
}