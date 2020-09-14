package main

import (
	"fmt"
	bnet "github.com/mrgreenturtle/bnetconnect/blizzardclient"
)

func main() {

	battletNetConnection := bnet.BnetClient{Client: bnet.CreateClient("a01c6ce324f8471f8996007eb0e352ca", "dhsvifXwB09ESsDAN0wsnhwbD3tC2zKg", "https://us.battle.net/oauth/token")}

	battletNetConnection.GeneratePvP("smilebomb", "kiljaeden")
	fmt.Printf("%+v\n", battletNetConnection.PvP[0].SeasonMatchStatistics)

}
