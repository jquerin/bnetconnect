package blizzardclient

import (
	ceq "bnetconnect/characterequipment"
	cp "bnetconnect/characterprofile"
	pvp "bnetconnect/characterpvp"
	pvps "bnetconnect/characterpvpsummary"
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
)

// BnetClient represents blizzard client api
type BnetClient struct {
	Client             *http.Client
	CharacterProfile   cp.CharacterProfile
	CharacterEquipment ceq.CharacterEquipment
	PvPSummary         pvps.CharacterPvpSummary
	PvP                [2]pvp.CharacterPvp
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
func (b *BnetClient) GenerateCharacterProfile(name string, realm string) {
	requestOut := fmt.Sprintf("https://us.api.blizzard.com/profile/wow/character/%s/%s?namespace=profile-us&locale=en_US", realm, name)
	resp, err := b.Client.Get(requestOut)
	if err != nil {
		fmt.Println(err)
	}
	profile := cp.CharacterProfile{}
	err = json.NewDecoder(resp.Body).Decode(&profile)
	b.CharacterProfile = profile
}

// GenerateCharacterEquipment method requests characterEquipment api
func (b *BnetClient) GenerateCharacterEquipment(name string, realm string) {
	requestOut := fmt.Sprintf("https://us.api.blizzard.com/profile/wow/character/%s/%s/equipment?namespace=profile-us&locale=en_US", realm, name)
	resp, err := b.Client.Get(requestOut)
	if err != nil {
		fmt.Println(err)
	}
	equipment := ceq.CharacterEquipment{}
	err = json.NewDecoder(resp.Body).Decode(&equipment)
	b.CharacterEquipment = equipment
}

func (b *BnetClient) GeneratePvP(name string, realm string) {
	requestOut := fmt.Sprintf("https://us.api.blizzard.com/profile/wow/character/%s/%s/pvp-bracket/3v3?namespace=profile-us&locale=en_US", realm, name)
	resp, err := b.Client.Get(requestOut)
	if err != nil {
		fmt.Println(err)
	}
	arena3v3 := pvp.CharacterPvp{}
	err = json.NewDecoder(resp.Body).Decode(&arena3v3)
	b.PvP[0] = arena3v3

	requestOut = fmt.Sprintf("https://us.api.blizzard.com/profile/wow/character/%s/%s/pvp-bracket/2v2?namespace=profile-us&locale=en_US", realm, name)
	resp, err = b.Client.Get(requestOut)
	if err != nil {
		fmt.Println(err)
	}
	arena2v2 := pvp.CharacterPvp{}
	err = json.NewDecoder(resp.Body).Decode(&arena2v2)
	b.PvP[1] = arena2v2

}

func (b *BnetClient) GeneratePvpSummary(name string, realm string) {
	requestOut := fmt.Sprintf("https://us.api.blizzard.com/profile/wow/character/%s/%s/pvp-summary?namespace=profile-us&locale=en_US", realm, name)
	resp, err := b.Client.Get(requestOut)
	if err != nil {
		fmt.Println(err)
	}
	pvpSummary := pvps.CharacterPvpSummary{}
	err = json.NewDecoder(resp.Body).Decode(&pvpSummary)
	b.PvPSummary = pvpSummary
}
