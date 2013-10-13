package main

import (
	"log"
	"time"
)

type Storage struct {
	ticker *time.Ticker    // Cleanup channel
	games map[string] Game // Make of GameIds to games
}

func NewStorage(ttl time.Duration) *Storage {
	s := new(Storage)
	s.ticker = time.NewTicker(ttl/2)
	go s.InitCleanup()
	return s
}

func (s *Storage) InitCleanup() {
	for {
		select {
		case <- s.ticker.C:
			s.cleanup()
		}
	}
}

func (s *Storage) cleanup() {
	log.Printf("Running cleanup")
}

func (s *Storage) GetGame(id string) (Game, error) {
	return s.games[id], nil
}

func (s *Storage) StartGame(id string) error {
	return nil
}