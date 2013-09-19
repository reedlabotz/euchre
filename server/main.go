package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"log"
	"strings"

	"code.google.com/p/go.net/websocket"
)

type Card string
type Deck [24]Card

const (
	ACE_OF_SPADES Card = "A♠"
	KING_OF_SPADES Card = "K♠"
	QUEEN_OF_SPADES Card = "Q♠"
	JACK_OF_SPADES Card = "J♠"
	TEN_OF_SPADES Card = "10♠"
	NINE_OF_SPADES Card = "9♠"

	ACE_OF_HEARTS Card = "A♥"
	KING_OF_HEARTS Card = "K♥"
	QUEEN_OF_HEARTS Card = "Q♥"
	JACK_OF_HEARTS Card = "J♥"
	TEN_OF_HEARTS Card = "10♥"
	NINE_OF_HEARTS Card = "9♥"

	ACE_OF_CLUBS Card = "A♦"
	KING_OF_CLUBS Card = "K♦"
	QUEEN_OF_CLUBS Card = "Q♦"
	JACK_OF_CLUBS Card = "J♦"
	TEN_OF_CLUBS Card = "10♦"
	NINE_OF_CLUBS Card = "9♦"

	ACE_OF_DIAMONDS Card = "A♣"
	KING_OF_DIAMONDS Card = "K♣"
	QUEEN_OF_DIAMONDS Card = "Q♣"
	JACK_OF_DIAMONDS Card = "J♣"
	TEN_OF_DIAMONDS Card = "10♣"
	NINE_OF_DIAMONDS Card = "9♣"
)
var DECK Deck = Deck{
	ACE_OF_SPADES,
	KING_OF_SPADES,
	QUEEN_OF_SPADES,
	JACK_OF_SPADES,
	TEN_OF_SPADES,
	NINE_OF_SPADES,
	ACE_OF_HEARTS,
	KING_OF_HEARTS,
	QUEEN_OF_HEARTS,
	JACK_OF_HEARTS,
	TEN_OF_HEARTS,
	NINE_OF_HEARTS,
	ACE_OF_CLUBS,
	KING_OF_CLUBS,
	QUEEN_OF_CLUBS,
	JACK_OF_CLUBS,
	TEN_OF_CLUBS,
	NINE_OF_CLUBS,
	ACE_OF_DIAMONDS,
	KING_OF_DIAMONDS,
	QUEEN_OF_DIAMONDS,
	JACK_OF_DIAMONDS,
	TEN_OF_DIAMONDS,
	NINE_OF_DIAMONDS,
}

const LOAD_SCREEN string = `
 _____ _   _ _____  _   _ ______ _____ 
|  ___| | | /  __ \| | | || ___ \  ___|
| |__ | | | | /  \/| |_| || |_/ / |__  
|  __|| | | | |    |  _  ||    /|  __| 
| |___| |_| | \__/\| | | || |\ \| |___ 
\____/ \___/ \____/\_| |_/\_| \_\____/`

func main() {
	fmt.Printf("%s\n\n", LOAD_SCREEN)
	setupHandler()
}

type Message struct {
	Text string
	Deck Deck
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
	return strings.Split(url.Path[1:],"/")[2], nil
}

func eucherServer(ws *websocket.Conn) {
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
		var m Message
		m.Text = message
		m.Deck = DECK
		websocket.JSON.Send(ws, &m)
		log.Printf("received: %s", message)
	}
}

func setupHandler() {
	http.Handle("/api/", websocket.Handler(eucherServer))
	err := http.ListenAndServe(":8020", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}