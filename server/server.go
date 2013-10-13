package main

import (
	"time"
	"code.google.com/p/go.net/websocket"
	"net/http"
	"log"
)

type Server struct {
	connections map[string] *websocket.Conn
	storage *Storage
}

func NewServer() *Server {
	s := new(Server)
	ttl := time.Duration(1)*time.Minute
	s.storage = NewStorage(ttl)
	return s
}

func (s *Server) Init() {
	http.HandleFunc("/api/game/new/", func(w http.ResponseWriter, r *http.Request) {
		s.HandleNew(w, r)
	})
	http.HandleFunc("/api/game/join/", func(w http.ResponseWriter, r *http.Request) {
		s.HandleJoin(w, r)
	})
	http.Handle("/api/game/play/", websocket.Handler(
		func(ws *websocket.Conn) {
			s.HandlePlay(ws) 
		}))
}

func (S *Server) Run() {
	err := http.ListenAndServe(":8020", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func (s *Server) HandlePlay(ws *websocket.Conn) {
	for {
		var message string
		err := websocket.Message.Receive(ws, &message)
		if err != nil {
			break
		}
		websocket.JSON.Send(ws, "a")
		log.Printf("received: %s", message)
	}
}

func (s *Server) HandleNew(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) HandleJoin(w http.ResponseWriter, r *http.Request) {

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
	SERVER_MOVE ServerMessageType = 1
)

type ServerMessage struct {
	Type ServerMessageType
	Move Move
	Game Game
	GameHash string
}