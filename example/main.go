package main

import (
	localnet "bnetconnect/blizzardclient"
	"fmt"
	bnet "github.com/mrgreenturtle/bnetconnect/blizzardclient"
	"os"
)

const tokenURL = "https://us.battle.net/oauth/token"

func main() {

	battletNetConnection := localnet.BnetClient{Client: bnet.CreateClient(os.Getenv("CLIENTID"), os.Getenv("CLIENTSECRET"), tokenURL)}
	profile := battletNetConnection.GenerateCharacterStats("smilebomb", "kiljaeden")
	fmt.Printf("%v", profile)
}
