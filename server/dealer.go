package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Hand struct {
	Cards []int
	Kitty map[string]int
}

type DealResponse struct {
	Dealer bool
	Hand   Hand
}

type Deal struct {
	Hands     map[string]Hand
	FreeHands []Hand
	Flip      int
}

const port = ":8020"

const gameTimeout = 15 * time.Second

const dealPath = "/deal/"

var deals = map[string]*Deal{}

/* Storage functions */

func makeNewDeal() *Deal {
	cards := rand.Perm(24)

	kitty := [3]int{cards[21], cards[22], cards[23]}

	for i := 0; i < 3; i++ {
		kitty[i] += rand.Intn(100) * 24
	}

	deal := new(Deal)
	for i := 0; i < 4; i++ {
		hand := new(Hand)
		hand.Kitty = make(map[string]int)
		for j := 0; j < 3; j++ {
			var k = 0
			if i < 3 {
				k = rand.Intn(kitty[j])
			} else {
				k = kitty[j]
			}

			kitty[j] -= k
			hand.Kitty[strconv.Itoa(j)] = k
			hand.Cards = cards[i*5 : i*5+5]
		}

		deal.FreeHands = append(deal.FreeHands, *hand)
		deal.Hands = make(map[string]Hand)
	}

	deal.Flip = cards[20]
	return deal
}

func deleteDeal(gameId string) {
	delete(deals, gameId)
}

func getDeal(gameId string) (*Deal, bool) {
	deal, exists := deals[gameId]

	if !exists {
		deal = makeNewDeal()
		deals[gameId] = deal
		log.Printf("# Game created: %s", gameId)
		time.AfterFunc(time.Duration(gameTimeout), func() {
			deleteDeal(gameId)
			log.Printf("# Game deleted: %s", gameId)
		})
	}

	return deal, !exists
}

/* Handler Functions */

func dealHandler(w http.ResponseWriter, r *http.Request) {

	urlParts := strings.Split(r.URL.Path, "/")

	if len(urlParts) < 4 {
		fmt.Fprintf(w, "Error")
		return
	}

	deal, dealer := getDeal(urlParts[2])

	hand, exists := deal.Hands[urlParts[3]]

	if !exists {
		if len(deal.FreeHands) == 0 {
			fmt.Fprintf(w, "Error")
			return
		}

		hand, deal.FreeHands = deal.FreeHands[len(deal.FreeHands)-1], deal.FreeHands[:len(deal.FreeHands)-1]

		deal.Hands[urlParts[3]] = hand
	} else {
		if len(deal.Hands) == 1 {
			dealer = true
		}
	}

	res := new(DealResponse)
	res.Hand = hand
	res.Dealer = dealer

	json, err := json.Marshal(res)

	if err != nil {
		fmt.Fprintf(w, "Error")
	}

	fmt.Fprintf(w, "%s", json)
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

/* main */

func main() {
	http.HandleFunc(dealPath, dealHandler)
	fmt.Printf("Server running on %s\n", port)
	http.ListenAndServe(port, Log(http.DefaultServeMux))
}
