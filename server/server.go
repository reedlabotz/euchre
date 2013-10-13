package main

import (
	"time"
	"code.google.com/p/go.net/websocket"
	"net/http"
	"net/url"
	"log"
	"errors"
	"net"
	"strings"
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
	http.Handle("/api/game/new/", websocket.Handler(
		func(ws *websocket.Conn) {
			s.HandlePlay(ws) 
		}))
	http.Handle("/api/game/join/", websocket.Handler(
		func(ws *websocket.Conn) {
			s.HandlePlay(ws) 
		}))
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
	var gameId, err = getGameId(ws.LocalAddr())
	if (err != nil) {
		log.Printf("Bad connection")
		ws.Close()
		return
	}
	log.Printf("Connection to game: %s",gameId)
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

func getGameId(addr net.Addr) (string, error) {
	var url, err = url.Parse(addr.String())
	if (err != nil) {
		return "", err
	}
	var parts = strings.Split(url.Path[1:], "/")
	if (len(parts) < 3) {
		return "", errors.New("Bad url path")
	}
	return strings.Split(url.Path[1:],"/")[3], nil
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