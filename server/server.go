package main

import (
	"time"
	"code.google.com/p/go.net/websocket"
	"net/http"
	"log"
	"github.com/gorilla/mux"
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
	r := mux.NewRouter()
	r.HandleFunc("/api/game/new/{GameId}/player/", func(w http.ResponseWriter, r *http.Request) {
		s.HandleNew(w, r)
	}).Methods("POST")
	r.HandleFunc("/api/game/join/{GameId}/player/", func(w http.ResponseWriter, r *http.Request) {
		s.HandleJoin(w, r)
	}).Methods("POST")
	r.HandleFunc("/api/game/join/{GameId}/player/{PlayerPublicKey}/{PlayerPrivateKey}/", func(w http.ResponseWriter, r *http.Request) {
		s.HandleJoin(w, r)
	}).Methods("POST")
	r.Handle("/api/game/play/{GameId}/player/{PlayerPublicKey}/{PlayerPrivateKey}/", websocket.Handler(
		func(ws *websocket.Conn) {
			s.HandlePlay(ws) 
		}))
	http.Handle("/", r)
}

func (S *Server) Run() {
	err := http.ListenAndServe(":8020", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func (s *Server) HandlePlay(ws *websocket.Conn) {
	vars := mux.Vars(ws.Request())
	log.Printf("Connection: %s [%s:%s]", vars["GameId"], vars["PlayerPublicKey"], vars["PlayerPrivateKey"])
	count := time.Duration(10)*time.Second
	ticker := time.NewTicker(count)
	go ping(ticker, ws)
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

func ping(t *time.Ticker, ws *websocket.Conn) {
	for {
		select {
		case <- t.C:
			websocket.Message.Send(ws, "Hi")
		}
	}
}

func (s *Server) HandleNew(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("New game: %s", vars["GameId"])
}

func (s *Server) HandleJoin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("Join game: %s [%s:%s]", vars["GameId"], vars["PlayerPublicKey"], vars["PlayerPrivateKey"])
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