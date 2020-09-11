package main

import (
	bnet "battle_net_connect/blizzardClient"
	player "battle_net_connect/player"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	alex := player.Player{"alex", "asdsd"}

	client := bnet.CreateClient(os.Getenv("CLIENTID"), os.Getenv("CLIENTSECRET"), os.Getenv("TOKENURL"))
	resp, err := client.Get("https://us.api.blizzard.com/profile/wow/character/area-52/bombsmile/equipment?namespace=profile-us&locale=en_US")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var f interface{}
	err = json.Unmarshal(body, &f)
	myData := f.(map[string]interface{})
	log.Println(myData["character"])

}
