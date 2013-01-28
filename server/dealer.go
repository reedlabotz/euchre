package main

import (
   "fmt"
   "net/http"
   "math/rand"
   "strings"
)

type Deal struct {

}

const port = ":8020"

const dealPath = "/deal/"


/* Handler Functions */

func dealHandler(w http.ResponseWriter, r *http.Request) {

   urlParts := strings.Split(r.URL.Path,"/")

   if len(urlParts) < 4 {
      fmt.Fprintf(w,"Error")
      return 
   }

   deck := rand.Perm(24)

   fmt.Fprintf(w,"URL Parts: %d\nDeck: %v",len(urlParts),deck)
}


/* main */

func main() {
   http.HandleFunc(dealPath,dealHandler)
   fmt.Printf("Server running on %s\n",port) 
   http.ListenAndServe(port,nil)
}