package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2/clientcredentials"
)

func main() {
	// Setup oauth2 config
	conf := &clientcredentials.Config{
		ClientID:     "a01c6ce324f8471f8996007eb0e352ca",
		ClientSecret: "dhsvifXwB09ESsDAN0wsnhwbD3tC2zKg",
		TokenURL:     "https://us.battle.net/oauth/token",
	}

	// Get a HTTP client
	client := conf.Client(context.Background())
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
