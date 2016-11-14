package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/stinkyfingers/socket/server/game"
	"golang.org/x/net/websocket"
	"gopkg.in/mgo.v2/bson"
)

// Client needs to:
// 1. Receive Lead Cards and Player Hand
// 2. Send Card (play)
// 3. Receive Played Cards
// 4. Send Card (vote)

var id = flag.String("id", "", "id--group id")
var player = flag.String("p", "", "-p=player id")

type Data struct {
	Message string
}

func main() {
	flag.Parse()
	if *id == "" || *player == "" {
		log.Fatal("-id is required and -p is required")
	}
	ws, err := websocket.Dial("ws://localhost:7000?id="+*id, "", "http://localhost")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			var g game.Game
			err := websocket.JSON.Receive(ws, &g)
			if err != nil {
				log.Print(err)
				break
			}

			log.Print("Dealer Cards: ", g.Round.DealerCards)
			//get self for logging
			for _, pl := range g.Players {
				if pl.ID.Hex() == *player {
					log.Print("Hand: ", pl.Hand)
				}
			}
			log.Print("Options: ", g.Round.Options)
			log.Print("Votes: ", g.Round.Votes)
			log.Print("Score: ", g.Round.Score)

		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		text := scanner.Text()

		arr := strings.Split(text, "=")
		if len(arr) < 2 {
			log.Print("Need type=value")
			break
		}
		switch arr[0] {
		case "play":
			play := game.Play{
				Card: game.Card{
					Phrase: arr[1],
				},
				Player: game.Player{
					ID: bson.ObjectIdHex(*player),
				},
				PlayType: "play",
			}
			websocket.JSON.Send(ws, play)

		case "vote":
			vote := game.Play{
				Card: game.Card{
					Phrase: arr[1],
				},
				Player: game.Player{
					ID: bson.ObjectIdHex(*player),
				},
				PlayType: "vote",
			}
			websocket.JSON.Send(ws, vote)

		default:
			log.Print("type of response is not allowed")
			break
		}

	}

}
