package blizzardclient

import (
	"context"
	"encoding/json"
	"fmt"
	ceq "github.com/mrgreenturtle/bnetconnect/characterequipment"
	cp "github.com/mrgreenturtle/bnetconnect/characterprofile"
	pvp "github.com/mrgreenturtle/bnetconnect/characterpvp"
	pvps "github.com/mrgreenturtle/bnetconnect/characterpvpsummary"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
)

// BnetClient represents blizzard client api
type BnetClient struct {
	Client *http.Client
}

// CreateClient creates http client to send out requests
func CreateClient(clientID string, clientSecret string, tokenURL string) *http.Client {
	conf := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     tokenURL,
	}
	client := conf.Client(context.Background())
	return client
}

// GenerateCharacterProfile method requests characterProfile api
func (b *BnetClient) GenerateCharacterProfile(name string, realm string) cp.CharacterProfile {
	requestOut := fmt.Sprintf("https://us.api.blizzard.com/profile/wow/character/%s/%s?namespace=profile-us&locale=en_US", realm, name)
	resp, err := b.Client.Get(requestOut)
	if err != nil {
		fmt.Println(err)
	}
	profile := cp.CharacterProfile{}
	err = json.NewDecoder(resp.Body).Decode(&profile)
	return profile
}

// GenerateCharacterEquipment method requests characterEquipment api
func (b *BnetClient) GenerateCharacterEquipment(name string, realm string) ceq.CharacterEquipment {
	requestOut := fmt.Sprintf("https://us.api.blizzard.com/profile/wow/character/%s/%s/equipment?namespace=profile-us&locale=en_US", realm, name)
	resp, err := b.Client.Get(requestOut)
	if err != nil {
		fmt.Println(err)
	}
	equipment := ceq.CharacterEquipment{}
	err = json.NewDecoder(resp.Body).Decode(&equipment)
	return equipment
}

func (b *BnetClient) GeneratePvP(name string, realm string) [2]pvp.CharacterPvp {
	requestOut := fmt.Sprintf("https://us.api.blizzard.com/profile/wow/character/%s/%s/pvp-bracket/3v3?namespace=profile-us&locale=en_US", realm, name)
	resp, err := b.Client.Get(requestOut)
	if err != nil {
		fmt.Println(err)
	}

	returnPvp := [2]pvp.CharacterPvp{}
	arena3v3 := pvp.CharacterPvp{}
	err = json.NewDecoder(resp.Body).Decode(&arena3v3)
	returnPvp[1] = arena3v3

	requestOut = fmt.Sprintf("https://us.api.blizzard.com/profile/wow/character/%s/%s/pvp-bracket/2v2?namespace=profile-us&locale=en_US", realm, name)
	resp, err = b.Client.Get(requestOut)
	if err != nil {
		fmt.Println(err)
	}
	arena2v2 := pvp.CharacterPvp{}
	err = json.NewDecoder(resp.Body).Decode(&arena2v2)
	returnPvp[0] = arena2v2
	return returnPvp
}

func (b *BnetClient) GeneratePvpSummary(name string, realm string) pvps.CharacterPvpSummary {
	requestOut := fmt.Sprintf("https://us.api.blizzard.com/profile/wow/character/%s/%s/pvp-summary?namespace=profile-us&locale=en_US", realm, name)
	resp, err := b.Client.Get(requestOut)
	if err != nil {
		fmt.Println(err)
	}
	pvpSummary := pvps.CharacterPvpSummary{}
	err = json.NewDecoder(resp.Body).Decode(&pvpSummary)
	return pvpSummary
}
