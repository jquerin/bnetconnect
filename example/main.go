package main

import (
	"fmt"
	bnet "github.com/mrgreenturtle/bnetconnect/blizzardclient"
	"os"
)

const tokenURL = "https://us.battle.net/oauth/token"

func main() {

	battletNetConnection := bnet.BnetClient{Client: bnet.CreateClient(os.Getenv("CLIENTID"), os.Getenv("CLIENTSECRET"), tokenURL)}
	profile := battletNetConnection.GeneratePvP("smilebomb", "kiljaeden")
	fmt.Printf("%v", profile)
}
