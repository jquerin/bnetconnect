package main

import (
	bnet "bnetconnect/blizzardClient"
	"fmt"
)

func main() {

	battletNetConnection := bnet.BnetClient{Client: bnet.CreateClient("a01c6ce324f8471f8996007eb0e352ca", "dhsvifXwB09ESsDAN0wsnhwbD3tC2zKg", "https://us.battle.net/oauth/token")}

	battletNetConnection.GenerateCharacterProfile("lunatara", "tichondrius")
	fmt.Printf("%v\n", battletNetConnection.CharacterProfile)

}
