package main

import (
   "fmt"
   "net/http"
   "strings"
   "math/rand"
   "encoding/json"
   "strconv"
   "time"
)

type Hand struct {
   Cards []int
   Kitty map[string] int
}

type DealResponse struct {
   Dealer bool
   Hand Hand
}

type Deal struct {
   Hands map[string]Hand
   FreeHands []Hand
   Flip int
}

const port = ":8020"

const dealPath = "/deal/"

var deals = map[string] *Deal {}


/* Storage functions */

func makeNewDeal() *Deal {
   cards := rand.Perm(24)

   kitty := [3]int{cards[21],cards[22],cards[23]}

   for i := 0; i < 3; i++ {
      kitty[i] += rand.Intn(100)*24
   }

   deal := new(Deal)
   for i := 0; i < 4; i++ {
      hand := new(Hand)
      hand.Kitty = make(map[string]int)
      for j := 0; j < 3; j++ {
         var k = 0
         if i < 3 {
            k = rand.Intn(kitty[j])
         }else{
            k = kitty[j]
         }
         
         kitty[j] -= k
         hand.Kitty[strconv.Itoa(j)] = k
         hand.Cards = cards[i*5:i*5+5]
      }

      deal.FreeHands = append(deal.FreeHands, *hand)
      deal.Hands = make(map[string]Hand)
   }

   deal.Flip = cards[20]
   return deal
}

func getDeal(gameId string) (*Deal, bool) {
   deal, exists := deals[gameId]

   if !exists {
      deal = makeNewDeal()
      deals[gameId] = deal
      time.AfterFunc(time.Duration(30*time.Second),func(){
         fmt.Println(deals)
         delete(deals,gameId)
         fmt.Println(deals)
      })
   }

   return deal, !exists
}


/* Handler Functions */

func dealHandler(w http.ResponseWriter, r *http.Request) {

   urlParts := strings.Split(r.URL.Path,"/")

   if len(urlParts) < 4 {
      fmt.Fprintf(w,"Error")
      return 
   }

   deal, dealer := getDeal(urlParts[2])

   hand, exists := deal.Hands[urlParts[3]]

   if !exists {
      if len(deal.FreeHands) == 0 {
         fmt.Fprintf(w,"Error")
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
      fmt.Fprintf(w,"Error")
   }

   fmt.Fprintf(w,"%s",json)
}


/* main */

func main() {
   http.HandleFunc(dealPath,dealHandler)
   fmt.Printf("Server running on %s\n",port) 
   http.ListenAndServe(port,nil)
}